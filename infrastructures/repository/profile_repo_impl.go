package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type ProfileRepoImpl struct {
	db *sql.DB
	tx *sql.Tx
}

func NewProfileRepoImpl(db *sql.DB) repository.ProfileRepo {
	return &ProfileRepoImpl{
		db: db,
	}
}

func (repo *ProfileRepoImpl) scanRow(row *sql.Row) (*model.Profile, error) {
	var profile model.Profile

	if err := row.Scan(
		&profile.ProfileID,
		&profile.UserID,
		&profile.Quote,
		&profile.CreatedAt,
		&profile.CreatedBy,
		&profile.UpdatedAt,
		&profile.UpdatedBy,
		&profile.DeletedAt,
		&profile.DeletedBy,
	); err != nil {
		log.Err(err).Msg(exception.LogErrScanning)
		return nil, err
	}
	return &profile, nil
}

func (repo *ProfileRepoImpl) GetProfileByID(ctx context.Context, id string) (*model.Profile, error) {
	query := "SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by " +
		"FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL"

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)

	profile, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (repo *ProfileRepoImpl) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	query := "SELECT id, user_id, quotes, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by " +
		"FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, userID)

	profile, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (repo *ProfileRepoImpl) StoreProfile(ctx context.Context, entity model.Profile) (model.Profile, error) {
	query := "SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1)"
	var exists bool

	querySTMT, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return model.Profile{}, err
	}
	if err = querySTMT.QueryRowContext(ctx, entity.UserID).Scan(&exists); err != nil {
		log.Err(err).Msg(exception.LogErrQuery)
		return model.Profile{}, err
	}

	if exists {
		return model.Profile{}, exception.Err400ProfileAlvailable
	}

	// insert proses
	query = "INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	execSTMT, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return model.Profile{}, err
	}

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.ProfileID,
		entity.UserID,
		entity.Quote,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.UpdatedAt,
	); err != nil {
		log.Err(err).Msg(exception.LogErrExec)
		return model.Profile{}, err
	}

	return entity, nil
}

func (repo *ProfileRepoImpl) UpdateProfile(ctx context.Context, entity model.Profile) (*model.Profile, error) {
	query := "UPDATE dueit.m_profiles SET quotes = $1, updated_by = $2, updated_at = $3 WHERE user_id = $4 AND id = $5 AND deleted_at IS NULL"

	stmt, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}

	if _, err = stmt.ExecContext(
		ctx,
		entity.Quote,
		entity.ProfileID,
		entity.UpdatedAt,
		entity.UserID,
		entity.ProfileID,
	); err != nil {
		log.Err(err).Msg(exception.LogErrExec)
		return nil, err
	}

	return &entity, nil
}

func (repo *ProfileRepoImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) error {
	tx, err := repo.db.BeginTx(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exception.LogErrTxStart)
		return err
	}

	repo.tx = tx
	return nil
}

func (repo *ProfileRepoImpl) GetTx() *sql.Tx {
	if repo.tx != nil {
		return repo.tx
	}
	return nil
}

func (repo *ProfileRepoImpl) Commit() error {
	if repo.tx != nil {
		err := repo.tx.Commit()
		if err != nil {
			log.Err(err).Msg(exception.LogErrTxCommit)
			return err
		}
		return nil
	}
	return exception.Err500TxNil
}

func (repo *ProfileRepoImpl) Rollback() error {
	err := repo.tx.Rollback()
	if err != nil {
		log.Err(err).Msg(exception.LogErrTxRollback)
		return err
	}

	return nil
}

func (repo *ProfileRepoImpl) CallTx(tx *sql.Tx) error {
	if tx != nil {
		repo.tx = tx
	}
	return exception.Err500TxNil
}
