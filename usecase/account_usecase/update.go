package account_usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/infra"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (a *AccountUsecaseImpl) UpdateAccount(ctx context.Context, req *usecase.RequestUpdateAccount) (*usecase.ResponseUser, *usecase.ResponseProfile, error) {
	profile, err := a.profileRepo.GetByID(ctx, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("profile id tidak di temukan")
			return nil, nil, usecase.ProfileNotFound
		}

		return nil, nil, err
	}

	user, err := a.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Info().Msg("user tidak di temukan")
			return nil, nil, usecase.UserNotFound
		}

		return nil, nil, err
	}

	if profile.UserID != user.ID {
		log.Info().Msg("profile user id tidak sama dengan request id")
		return nil, nil, usecase.ProfileUserIDAndReqUserIDNotMatch
	}

	exist, err := a.userRepo.CheckPhoneNumberExists(ctx, req.UserID, req.PhoneNumber)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		return nil, nil, usecase.PhoneNumberIsExist
	}

	oldImage := user.Image
	newImageName := user.Image
	email := user.Email
	reqImageCondition := req.Image != nil && req.Image.Size > 0
	delImageCondition := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google") && reqImageCondition

	if reqImageCondition {
		fileExt := filepath.Ext(req.Image.Filename)
		newImageName = a.minioRepo.GenerateFileName(fileExt, "user-images/public/")
		newImageName = fmt.Sprintf("/%s/%s", infra.MinIoBucket, newImageName)
	}

	profile, user = req.ToModel(newImageName)

	err = a.profileRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
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

	userResp := &usecase.ResponseUser{
		ID:              user.ID,
		FullName:        user.FullName,
		Gender:          user.Gender,
		Image:           user.Image,
		Username:        user.Username,
		Email:           user.Email,
		EmailFormat:     usecase.EmailFormat(email),
		PhoneNumber:     repository.GetNullString(user.PhoneNumber),
		EmailVerifiedAt: user.EmailVerifiedAt,
	}
	profileResp := &usecase.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     repository.GetNullString(profile.Quote),
		Profesi:   repository.GetNullString(profile.Profesi),
	}
	return userResp, profileResp, nil
}
