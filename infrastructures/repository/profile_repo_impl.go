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
	uow repository.UnitOfWork
}

func NewProfileRepoImpl(uow repository.UnitOfWork) repository.ProfileRepo {
	return &ProfileRepoImpl{
		uow: uow,
	}
}

func (repo *ProfileRepoImpl) UoW() repository.UnitOfWork {
	return repo.uow
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
	tx, err := repo.uow.GetTx()
	if err != nil {
		return model.Profile{}, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return model.Profile{}, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

	if err = querySTMT.QueryRowContext(ctx, entity.UserID).Scan(&exists); err != nil {
		log.Err(err).Msg(exception.LogErrQuery)
		return model.Profile{}, err
	}

	if exists {
		return model.Profile{}, exception.Err400ProfileAlvailable
	}

	// insert proses
	query = "INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return model.Profile{}, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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
	tx, err := repo.uow.GetTx()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrSTMT)
		return nil, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(exception.LogErrCloseStmt)
		}
	}()

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
