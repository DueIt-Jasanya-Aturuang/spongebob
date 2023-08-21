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
)

type ProfileUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	tsx         repository.SqlTransactionRepo
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo repository.ProfileRepo,
	tsx repository.SqlTransactionRepo,
	timeout time.Duration,
) usecase.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		tsx:         tsx,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileById(c context.Context, id string) (*dto.ProfileResp, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var resp dto.ProfileResp
	res, err := u.profileRepo.GetProfileById(ctx, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		res, err = u.profileRepo.GetProfileByUserId(ctx, id)
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

func (u *ProfileUsecaseImpl) storeProfile(c context.Context, userId string) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var profile model.Profile
	profile.UserId = userId
	profile = profile.DefaultValue()

	err := u.tsx.Transaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		profileRes, err := u.profileRepo.StoreProfile(ctx, tx, profile)
		if err != nil {
			return err
		}
		profile = profileRes
		return nil
	})
	return &profile, err
}
