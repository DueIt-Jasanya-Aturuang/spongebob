package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
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

func (u *ProfileUsecaseImpl) GetProfileByID(c context.Context, id string) (*dto.ProfileResp, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var resp dto.ProfileResp
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
			return &resp, nil
		}
	}

	resp = res.ToResp()
	return &resp, nil
}

func (u *ProfileUsecaseImpl) storeProfile(c context.Context, userID string) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var profile model.Profile
	profile.UserID = userID
	profile = profile.DefaultValue()

	err := u.profileRepo.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	defer func() {
		if err != nil {
			log.Info().Msg(exception.LogInfoTxRollback)
			if rbErr := u.profileRepo.Rollback(); rbErr != nil {
				log.Err(rbErr).Msg(exception.LogErrTxRollback)
				err = fmt.Errorf("error : %v || rbError : %v", err, rbErr)
			}
			return
		}
		log.Info().Msg(exception.LogInfoTxCommit)
		if cmErr := u.profileRepo.Commit(); cmErr != nil {
			log.Err(cmErr).Msg(exception.LogErrTxCommit)
			err = fmt.Errorf("error commit: %v", cmErr)
		}
	}()

	if err != nil {
		return nil, err
	}

	profileRes, err := u.profileRepo.StoreProfile(ctx, profile)
	if err != nil {
		return nil, err
	}
	profile = profileRes
	return &profile, err
}
