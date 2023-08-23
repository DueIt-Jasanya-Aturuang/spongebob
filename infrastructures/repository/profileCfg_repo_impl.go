package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type ProfileCfgRepoImpl struct {
	db *sql.DB
	tx *sql.Tx
}

func NewProfileCfgRepoImpl(db *sql.DB) repository.ProfileCfgRepo {
	return &ProfileCfgRepoImpl{
		db: db,
	}
}

func (repo *ProfileCfgRepoImpl) scanRow(row *sql.Row) (*model.ProfileCfg, error) {
	var profileCfg model.ProfileCfg

	if err := row.Scan(
		&profileCfg.ID,
		&profileCfg.ProfileID,
		&profileCfg.ConfigName,
		&profileCfg.ConfigValue,
		&profileCfg.Status,
		&profileCfg.CreatedAt,
		&profileCfg.CreatedBy,
		&profileCfg.UpdatedAt,
		&profileCfg.UpdatedBy,
		&profileCfg.DeletedAt,
		&profileCfg.DeletedBy,
	); err != nil {
		log.Err(err).Msg(exception.LogErrScanning)
		return nil, err
	}
	return &profileCfg, nil
}

func (repo *ProfileCfgRepoImpl) scanRows(rows *sql.Rows) (*[]model.ProfileCfg, error) {
	var profileCfgs []model.ProfileCfg

	for rows.Next() {
		var profileCfg model.ProfileCfg
		if err := rows.Scan(
			&profileCfg.ID,
			&profileCfg.ProfileID,
			&profileCfg.ConfigName,
			&profileCfg.ConfigValue,
			&profileCfg.Status,
			&profileCfg.CreatedAt,
			&profileCfg.CreatedBy,
			&profileCfg.UpdatedAt,
			&profileCfg.UpdatedBy,
			&profileCfg.DeletedAt,
			&profileCfg.DeletedBy,
		); err != nil {
			log.Err(err).Msg(exception.LogErrScanning)
			return nil, err
		}
		profileCfgs = append(profileCfgs, profileCfg)
	}
	return &profileCfgs, nil
}

func (repo *ProfileCfgRepoImpl) GetProfileCfgByID(ctx context.Context, id string) (*model.ProfileCfg, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
              FROM dueit.m_user_config WHERE profile_id = $1 OR id = $2`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id, id)

	profileCfg, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return profileCfg, nil
}

func (repo *ProfileCfgRepoImpl) GetProfileCfgByScheduler(ctx context.Context, profileCfgScheduler dto.ProfileCfgScheduler) (*[]model.ProfileCfg, error) {
	query := `SELECT id, profile_id, config_name, config_value, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
              FROM dueit.m_user_config WHERE (config_value->>'config_time_notify')::time >= $1::time AND config_value->'days' ? $2`
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, profileCfgScheduler.Time, profileCfgScheduler.Day)
	if err != nil {
		log.Err(err).Msg(exception.LogErrQuery)
	}

	profileCfgs, err := repo.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return profileCfgs, nil
}

func (repo *ProfileCfgRepoImpl) StoreProfileCfg(ctx context.Context, profileCfg model.ProfileCfg) error {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2)`
	var exists bool
	querySTMT, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}
	if err = querySTMT.QueryRowContext(ctx, profileCfg.ProfileID, profileCfg.ConfigName).Scan(&exists); err != nil {
		log.Err(err).Msg(exception.LogErrQuery)
		return err
	}
	if exists {
		return exception.Err400ProfileConfigAlvailable
	}

	// process insert
	query = `INSERT INTO dueit.m_user_config (id, profile_id, config_name, config_value, status, created_at, created_by, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	execSTMT, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}

	_, err = execSTMT.ExecContext(
		ctx,
		profileCfg.ID,
		profileCfg.ProfileID,
		profileCfg.ConfigName,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.CreatedAt,
		profileCfg.CreatedBy,
		profileCfg.UpdatedAt,
	)
	return err
}

func (repo *ProfileCfgRepoImpl) UpdateProfileCfg(ctx context.Context, profileCfg model.ProfileCfg) error {
	query := `UPDATE dueit.m_user_config SET config_value = $1, status = $2, updated_at = $3, updated_by = $4 WHERE id = $5`
	stmt, err := repo.tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}

	_, err = stmt.ExecContext(
		ctx,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.UpdatedAt,
		profileCfg.UpdatedBy,
		profileCfg.ID,
	)
	return err
}

func (repo *ProfileCfgRepoImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) error {
	tx, err := repo.db.BeginTx(ctx, opts)
	if err != nil {
		log.Err(err).Msg(exception.LogErrTxStart)
		return err
	}

	repo.tx = tx
	return nil
}

func (repo *ProfileCfgRepoImpl) GetTx() *sql.Tx {
	if repo.tx != nil {
		return repo.tx
	}
	return nil
}

func (repo *ProfileCfgRepoImpl) Commit() error {
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

func (repo *ProfileCfgRepoImpl) Rollback() error {
	err := repo.tx.Rollback()
	if err != nil {
		log.Err(err).Msg(exception.LogErrTxRollback)
		return err
	}

	return nil
}

func (repo *ProfileCfgRepoImpl) CallTx(tx *sql.Tx) error {
	if tx != nil {
		repo.tx = tx
	}
	return exception.Err500TxNil
}
