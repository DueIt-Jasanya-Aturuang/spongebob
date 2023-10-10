package account_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (a *AccountUsecaseImpl) GetProfileByUserID(ctx context.Context, userID string) (*usecase.ResponseProfile, error) {
	profile, err := a.profileRepo.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ProfileNotFound
		}
		return nil, err
	}

	resp := &usecase.ResponseProfile{
		ProfileID: profile.ProfileID,
		Quote:     repository.GetNullString(profile.Quote),
		Profesi:   repository.GetNullString(profile.Profesi),
	}
	return resp, nil
}
