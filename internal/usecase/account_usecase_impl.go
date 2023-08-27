package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
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

func (u *AccountUsecaseImpl) UpdateAccount(c context.Context, req dto.UpdateAccountReq) (userResp *dto.UserResp, profileResp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	if profile.UserID != user.ID {
		return nil, nil, exception.Err401Unauthorization
	}

	oldImage := user.Image
	email := user.Email
	reqImageBool := req.Image != nil && req.Image.Size > 0
	delImageBool := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google")

	var newImageName string
	if reqImageBool {
		newImageName = u.minioClient.GenerateFileName(req.Image, "user-images/public/", "")
		user.Image = fmt.Sprintf("/%s/%s", config.MinIoBucket, newImageName)
	}

	profileConv, userConv := dtoconv.UpdateAccountToModel(req, user.Image)

	profileRepoUOW := u.profileRepo.UoW()
	err = profileRepoUOW.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if errEndTx := profileRepoUOW.EndTx(err); errEndTx != nil {
			err = errEndTx
			profileResp = nil
			userResp = nil
		}
	}()

	profile, err = u.profileRepo.UpdateProfile(ctx, profileConv)
	if err != nil {
		return nil, nil, err
	}

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

	if reqImageBool {
		err = u.minioClient.UploadFile(ctx, req.Image, newImageName, config.MinIoBucket)
		if err != nil {
			return nil, nil, err
		}

		if delImageBool {
			oldImageArr := strings.Split(oldImage, "/")
			newImageName = fmt.Sprintf("/%s/%s/%s", oldImageArr[2], oldImageArr[3], oldImageArr[4])
			err = u.minioClient.DeleteFile(ctx, newImageName, config.MinIoBucket)
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
