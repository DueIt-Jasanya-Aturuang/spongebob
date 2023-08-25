package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/dtoconv"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/format"
)

type AccountUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	userRepo    repository.UserRepo
	minioClient repository.MinioRepo
	ctxTimeout  time.Duration
}

func NewAccountUsecaseImpl(
	profileRepo repository.ProfileRepo,
	userRepo repository.UserRepo,
	minioClient repository.MinioRepo,
	ctxTimeout time.Duration,
) usecase.AccountUsecase {
	return &AccountUsecaseImpl{
		profileRepo: profileRepo,
		userRepo:    userRepo,
		minioClient: minioClient,
		ctxTimeout:  ctxTimeout,
	}
}

func (u *AccountUsecaseImpl) AccountUpdate(c context.Context, req dto.UpdateAccountReq) (userResp *dto.UserResp, profileResp *dto.ProfileResp, err error) {
	// set timeout process
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	// get profile by user id request
	profile, err := u.profileRepo.GetProfileByUserID(ctx, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	// get user by id request
	user, err := u.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	// condition req image and oldimage
	oldImage := user.Image
	email := user.Email
	reqImageBool := req.Image != nil && req.Image.Size > 0
	delImageBool := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google")

	// convert dto to model
	profileConv, userConv := dtoconv.UpdateAccountToModel(req, profile.ProfileID, user.Image)

	// declar profile repo unit of work
	profileRepoUOW := u.profileRepo.UoW()

	// start tx from profile repo
	err = profileRepoUOW.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		errEndTx := profileRepoUOW.EndTx(err)
		if errEndTx != nil {
			err = errEndTx
			profileResp = nil
			userResp = nil
		}
	}()

	// update profile repo process
	profile, err = u.profileRepo.UpdateProfile(ctx, profileConv)
	if err != nil {
		return nil, nil, err
	}

	// call tx from profile repo to user repo
	userRepoUOW := u.userRepo.UoW()
	txProfileRepoUOW, err := profileRepoUOW.GetTx()
	if err != nil {
		return nil, nil, err
	}

	err = userRepoUOW.CallTx(txProfileRepoUOW)
	if err != nil {
		return nil, nil, err
	}
	user, err = u.userRepo.UpdateUser(ctx, userConv)
	if err != nil {
		return nil, nil, err
	}

	filename := u.minioClient.GenerateFileName(req.Image, "user-images/public/", "")
	if reqImageBool {
		user.Image = fmt.Sprintf("/%s/%s", config.MinIoBucket, filename)
	}

	if reqImageBool {
		err = u.minioClient.UploadFile(ctx, req.Image, filename, config.MinIoBucket)
		if err != nil {
			return nil, nil, err
		}

		if delImageBool {
			oldImageArr := strings.Split(oldImage, "/")
			filename = fmt.Sprintf("/%s/%s/%s", oldImageArr[2], oldImageArr[3], oldImageArr[4])
			err = u.minioClient.DeleteFile(ctx, filename, config.MinIoBucket)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	emailFormat, err := format.EmailFormat(email)
	if err != nil {
		return nil, nil, err
	}
	userResp = user.ToResp(emailFormat)
	profileResp = profile.ToResp()
	return userResp, profileResp, nil
}
