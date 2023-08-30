package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"strings"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
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

func (u *AccountUsecaseImpl) UpdateAccount(c context.Context, req *dto.UpdateAccountReq) (userResp *dto.UserResp, profileResp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	profile, err := u.profileRepo.GetProfileByID(ctx, req.ProfileID)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, model.ErrUnauthorization
		}
		return nil, nil, err
	}

	if profile.UserID != user.ID {
		return nil, nil, model.ErrUnauthorization
	}

	oldImage := user.Image
	email := user.Email
	reqImageCondition := req.Image != nil && req.Image.Size > 0
	delImageCondition := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google")

	var newImageName string
	if reqImageCondition {
		newImageName = u.minioClient.GenerateFileName(req.Image, "user-images/public/", "")
		user.Image = fmt.Sprintf("/%s/%s", config.MinIoBucket, newImageName)
	}

	phoneNumberExists, err := u.userRepo.CheckPhoneNumberExists(ctx, user.ID, req.PhoneNumber)
	if err != nil {
		return nil, nil, err
	}
	if phoneNumberExists {
		return nil, nil, response.Err409(map[string][]string{
			"phone_number": {
				message.PhoneNumberIsAvailable,
			},
		}, err)
	}

	profileConv, userConv := helpers.UpdateAccountToModel(req, user.Image)

	err = u.profileRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if errEndTx := u.profileRepo.EndTx(err); errEndTx != nil {
			err = errEndTx
			profileResp = nil
			userResp = nil
		}
	}()

	profile, err = u.profileRepo.UpdateProfile(ctx, profileConv)
	if err != nil {
		return nil, nil, err
	}

	user, err = u.userRepo.UpdateUser(ctx, userConv)
	if err != nil {
		return nil, nil, err
	}

	if reqImageCondition {
		err = u.minioClient.UploadFile(ctx, req.Image, newImageName, config.MinIoBucket)
		if err != nil {
			return nil, nil, err
		}

		if delImageCondition {
			oldImageArr := strings.Split(oldImage, "/")
			newImageName = fmt.Sprintf("/%s/%s/%s", oldImageArr[2], oldImageArr[3], oldImageArr[4])
			err = u.minioClient.DeleteFile(ctx, newImageName, config.MinIoBucket)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	emailFormat := helpers.EmailFormat(email)

	userResp = user.ToResp(emailFormat)
	profileResp = profile.ToResp()

	return userResp, profileResp, nil
}
