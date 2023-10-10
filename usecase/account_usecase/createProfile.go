package account_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (a *AccountUsecaseImpl) CreateProfile(ctx context.Context, userID string) (*usecase.ResponseProfile, error) {
	_, err := a.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.UserNotFound
		}
		return nil, err
	}

	profileResp, err := a.GetProfileByUserID(ctx, userID)
	if err != nil {
		if !errors.Is(err, usecase.ProfileNotFound) {
			return nil, err
		}
	}

	if errors.Is(err, usecase.ProfileNotFound) {
		return profileResp, nil
	}

	profile := repository.DefaultValueProfile(userID)
	err = a.profileRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		exist, err := a.profileRepo.Create(ctx, profile)
		if err != nil {
			return err
		}

		if exist {
			log.Warn().Msgf("user sudah memiliki profile tapi create lagi | data : %v", profile)
		}
		return nil
	})

	resp := &usecase.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     repository.GetNullString(profile.Quote),
		Profesi:   repository.GetNullString(profile.Profesi),
	}
	return resp, nil
}
