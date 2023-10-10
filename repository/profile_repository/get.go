package profile_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileRepositoryImpl) GetByID(ctx context.Context, id string) (*repository.Profile, error) {
	query := `SELECT id, user_id, quotes, profesi, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL`

	db, err := p.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profile repository.Profile
	if err = stmt.QueryRowContext(ctx, id).Scan(
		&profile.ProfileID,
		&profile.UserID,
		&profile.Quote,
		&profile.Profesi,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &profile, nil
}

func (p *ProfileRepositoryImpl) GetByUserID(ctx context.Context, userID string) (*repository.Profile, error) {
	query := `SELECT id, user_id, quotes, profesi, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL`

	db, err := p.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var profile repository.Profile
	if err = stmt.QueryRowContext(ctx, userID).Scan(
		&profile.ProfileID,
		&profile.UserID,
		&profile.Quote,
		&profile.Profesi,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &profile, nil
}
