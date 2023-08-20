package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	domainprofile "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/db"
)

type ProfileUsecaseImpl struct {
	profileRepo domainprofile.ProfileRepo
	db          db.SQL
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo domainprofile.ProfileRepo,
	db db.SQL,
	timeout time.Duration,
) domainprofile.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		db:          db,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileById(c context.Context, id string) (*domainprofile.ProfileResp, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var resp domainprofile.ProfileResp
	res, err := u.profileRepo.GetProfileById(ctx, u.db.SqlDB(), id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		res, err = u.profileRepo.GetProfileByUserId(ctx, u.db.SqlDB(), id)
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

func (u *ProfileUsecaseImpl) storeProfile(c context.Context, userId string) (*domainprofile.Profile, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	var profile domainprofile.Profile
	profile.UserId = userId
	profile = profile.DefaultValue()

	err := u.db.Transaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		profileRes, err := u.profileRepo.StoreProfile(ctx, tx, profile)
		if err != nil {
			return err
		}
		profile = profileRes
		return nil
	})
	return &profile, err
}
