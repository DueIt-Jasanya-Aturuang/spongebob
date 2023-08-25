package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
	"github.com/rs/zerolog/log"
)

type ProfileUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo repository.ProfileRepo,
	timeout time.Duration,
) usecase.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileByID(c context.Context, id string) (resp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	res, err := u.profileRepo.GetProfileByID(ctx, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		res, err = u.profileRepo.GetProfileByUserID(ctx, id)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
			// store profile
			profile, err := u.storeProfile(ctx, id)
			if err != nil {
				return nil, err
			}

			resp = profile.ToResp()
			return resp, nil
		}
	}

	resp = res.ToResp()
	return resp, nil
}

func (u *ProfileUsecaseImpl) storeProfile(c context.Context, userID string) (profile *model.Profile, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile = profile.DefaultValue(userID)

	// declar profile repo unit of work
	profileRepoUOW := u.profileRepo.UoW()

	// start tx from profile repo
	err = profileRepoUOW.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		errEndTx := profileRepoUOW.EndTx(err)
		if errEndTx != nil {
			err = errEndTx
			profile = nil
		}
	}()
	profileRes, err := u.profileRepo.StoreProfile(ctx, *profile)
	if err != nil {
		return nil, err
	}
	profile = &profileRes
	log.Info().Msgf("%v", profile)
	return profile, err
}
