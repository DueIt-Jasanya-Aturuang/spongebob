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
	uow repository.UnitOfWork
}

func NewProfileCfgRepoImpl(uow repository.UnitOfWork) repository.ProfileCfgRepo {
	return &ProfileCfgRepoImpl{
		uow: uow,
	}
}

func (repo *ProfileCfgRepoImpl) UoW() repository.UnitOfWork {
	return repo.uow
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

	conn, err := repo.uow.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errConn := conn.Close(); errConn != nil {
			log.Err(errConn).Msg(exception.LogErrCloseConn)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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

	conn, err := repo.uow.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errConn := conn.Close(); errConn != nil {
			log.Err(errConn).Msg(exception.LogErrCloseConn)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

	rows, err := stmt.QueryContext(ctx, profileCfgScheduler.Time, profileCfgScheduler.Day)
	if err != nil {
		log.Err(err).Msg(exception.LogErrQuery)
	}
	defer func() {
		if errRows := rows.Close(); errRows != nil {
			log.Err(errRows).Msg(exception.LogErrCloseRows)
		}
	}()

	profileCfgs, err := repo.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return profileCfgs, nil
}

func (repo *ProfileCfgRepoImpl) StoreProfileCfg(ctx context.Context, profileCfg model.ProfileCfg) error {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2)`
	tx, err := repo.uow.GetTx()
	if err != nil {
		return err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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
	tx, err := repo.uow.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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
