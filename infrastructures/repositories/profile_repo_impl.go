package repositories

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type ProfileRepoImpl struct {
	DB *sql.DB
}

func NewProfileRepoImpl(db *sql.DB) repository.ProfileRepo {
	return &ProfileRepoImpl{
		DB: db,
	}
}

func (repo *ProfileRepoImpl) scanRow(row *sql.Row) (*model.Profile, error) {
	var profile model.Profile

	if err := row.Scan(
		&profile.ProfileId,
		&profile.UserId,
		&profile.Quote,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		log.Err(err).Msg(exceptions.LogErrScanning)
		return nil, err
	}
	return &profile, nil
}

func (repo *ProfileRepoImpl) GetProfileById(ctx context.Context, id string) (*model.Profile, error) {
	query := "SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL"
	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)

	profile, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (repo *ProfileRepoImpl) GetProfileByUserId(ctx context.Context, userId string) (*model.Profile, error) {
	query := "SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL"
	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, userId)

	profile, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (repo *ProfileRepoImpl) StoreProfile(ctx context.Context, tx *sql.Tx, entity model.Profile) (model.Profile, error) {
	query := "SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1)"
	var exists bool

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return model.Profile{}, err
	}
	if err = querySTMT.QueryRowContext(ctx, entity.UserId).Scan(&exists); err != nil {
		log.Err(err).Msg(exceptions.LogErrQuery)
		return model.Profile{}, err
	}
	if exists {
		return model.Profile{}, exceptions.ErrProfileAlvailable
	}

	// insert proses
	query = "INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return model.Profile{}, err
	}

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.ProfileId,
		entity.UserId,
		entity.Quote,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.UpdatedAt,
	); err != nil {
		log.Err(err).Msg(exceptions.LogErrExec)
		return model.Profile{}, err
	}

	return entity, nil
}

func (repo *ProfileRepoImpl) UpdateProfile(ctx context.Context, tx *sql.Tx, entity model.Profile) (*model.Profile, error) {
	query := "UPDATE dueit.m_profiles SET quotes = $1, updated_by = $2, updated_at = $3 WHERE user_id = $4 AND id = $5 AND deleted_at IS NULL"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	if _, err = stmt.ExecContext(
		ctx,
		entity.Quote,
		entity.ProfileId,
		entity.UpdatedAt,
		entity.UserId,
		entity.ProfileId,
	); err != nil {
		log.Err(err).Msg(exceptions.LogErrExec)
		return nil, err
	}

	return &entity, nil
}
