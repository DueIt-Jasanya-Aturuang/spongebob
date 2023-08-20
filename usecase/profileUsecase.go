package usecase

import (
	"context"
	"database/sql"
	"time"

	domainprofile "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/db"
)

type ProfileUsecaseImpl struct {
	profileRepo domainprofile.ProfileRepo
	db          *db.DBimpl
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo domainprofile.ProfileRepo,
	db *db.DBimpl,
	timeout time.Duration,
) domainprofile.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		db:          db,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileById(ctx context.Context, id string) (*domainprofile.ProfileResp, error) {
	return nil, nil
}

func (u *ProfileUsecaseImpl) StoreProfile(ctx context.Context, userId string) (*domainprofile.Profile, error) {
	profile := domainprofile.Profile{}

	err := u.db.Transaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		_, err := u.profileRepo.StoreProfile(ctx, tx, profile)
		if err != nil {
			return err
		}
		return nil
	})
	return &profile, err
}
