package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/helpers"
)

type AccountUsecaseImpl struct {
	profileRepo domain.ProfileRepo
	userRepo    domain.UserRepo
	minioRepo   domain.MinioRepo
}

func NewAccountUsecaseImpl(
	profileRepo domain.ProfileRepo,
	userRepo domain.UserRepo,
	minioRepo domain.MinioRepo,
) domain.AccountUsecase {
	return &AccountUsecaseImpl{
		profileRepo: profileRepo,
		userRepo:    userRepo,
		minioRepo:   minioRepo,
	}
}

func (a *AccountUsecaseImpl) UpdateAccount(ctx context.Context, req *domain.RequestUpdateAccount) (*domain.ResponseUser, *domain.ResponseProfile, error) {
	err := a.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer a.profileRepo.CloseConn()

	profile, err := a.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, nil, ProfileNotFound
		}

		return nil, nil, err
	}

	user, err := a.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("user tidak di temukan")
			return nil, nil, UserNotFound
		}

		return nil, nil, err
	}

	if profile.UserID != user.ID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, nil, ProfileUserIDAndReqUserIDNotMatch
	}

	exist, err := a.userRepo.CheckPhoneNumberExists(ctx, req.UserID, req.PhoneNumber)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		return nil, nil, PhoneNumberIsExist
	}

	oldImage := user.Image
	newImageName := user.Image
	email := user.Email
	reqImageCondition := req.Image != nil && req.Image.Size > 0
	delImageCondition := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google") && reqImageCondition

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		newImageName = a.minioRepo.GenerateFileName(fileExt, "user-images/public/")
	}

	profile, user = converter.UpdateAccountToModel(req, fmt.Sprintf("/%s/%s", infra.MinIoBucket, newImageName))

	err = a.profileRepo.StartTx(ctx, helpers.LevelReadCommitted(), func() error {
		err = a.profileRepo.Update(ctx, profile)
		if err != nil {
			return err
		}

		err = a.userRepo.Update(ctx, user)
		if err != nil {
			return err
		}

		if reqImageCondition {
			if err = a.minioRepo.UploadFile(ctx, req.Image, newImageName); err != nil {
				return err
			}

		}

		return nil
	})

	if delImageCondition {
		imageDelArr := strings.Split(oldImage, "/")
		if len(imageDelArr) == 4 {
			imageDel := fmt.Sprintf("/%s/%s/%s", imageDelArr[2], imageDelArr[3], imageDelArr[4])
			if err = a.minioRepo.DeleteFile(ctx, imageDel); err != nil {
				return nil, nil, err
			}
		}

	}

	if err != nil {
		return nil, nil, err
	}

	emailFormat := helpers.EmailFormat(email)

	userResp, profileResp := converter.UpdateAccountModelToResp(user, profile, emailFormat)
	return userResp, profileResp, nil
}

func (a *AccountUsecaseImpl) GetProfileByUserID(ctx context.Context, req *domain.RequestGetProfile) (*domain.ResponseProfile, error) {
	err := a.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer a.profileRepo.CloseConn()

	profile, err := a.profileRepo.GetByUserID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ProfileNotFound
		}
		return nil, err
	}

	resp := converter.ProfileModelToResp(profile)
	return resp, nil
}

func (a *AccountUsecaseImpl) CreateProfile(ctx context.Context, req *domain.RequestCreateProfile) (*domain.ResponseProfile, error) {
	err := a.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer a.profileRepo.CloseConn()

	_, err = a.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, UserNotFound
		}
		return nil, err
	}

	profile := converter.ProfileDefault(req.UserID)

	err = a.profileRepo.StartTx(ctx, helpers.LevelReadCommitted(), func() error {
		exist, err := a.profileRepo.Create(ctx, profile)
		if err != nil {
			return err
		}

		if exist {
			log.Warn().Msgf("user sudah memiliki profile tapi create lagi | data : %v", req)
		}
		return nil
	})

	resp := converter.ProfileModelToResp(profile)
	return resp, nil
}
