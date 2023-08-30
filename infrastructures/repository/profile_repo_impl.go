package repository

import (
	"context"
	"database/sql"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"github.com/rs/zerolog/log"
)

type ProfileRepoImpl struct {
	repository.UnitOfWork
}

func NewProfileRepoImpl(uow repository.UnitOfWork) repository.ProfileRepo {
	return &ProfileRepoImpl{
		UnitOfWork: uow,
	}
}

func (repo *ProfileRepoImpl) GetProfileByID(ctx context.Context, id string) (model.Profile, error) {
	query := `SELECT id, user_id, quotes, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE id = $1 AND deleted_at IS NULL`

	conn, err := repo.GetConn()
	if err != nil {
		return model.Profile{}, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.Profile{}, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	row := stmt.QueryRowContext(ctx, id)

	profile, err := repo.scanRow(row)
	if err != nil {
		return model.Profile{}, err
	}
	return profile, nil
}

// GetProfileByUserID get profile by user id
func (repo *ProfileRepoImpl) GetProfileByUserID(ctx context.Context, userID string) (model.Profile, error) {
	query := `SELECT id, user_id, quotes, created_at, created_by, 
       				 updated_at, updated_by, deleted_at, deleted_by 
			  FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL`

	conn, err := repo.GetConn()
	if err != nil {
		return model.Profile{}, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.Profile{}, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	row := stmt.QueryRowContext(ctx, userID)

	profile, err := repo.scanRow(row)
	if err != nil {
		return model.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepoImpl) StoreProfile(ctx context.Context, profile model.Profile) (model.Profile, error) {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL)`
	var exists bool

	tx, err := repo.GetTx()
	if err != nil {
		return model.Profile{}, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.Profile{}, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if err = querySTMT.QueryRowContext(ctx, profile.UserID).Scan(&exists); err != nil {
		log.Err(err).Msg(message.ErrQueryRowDB)
		return model.Profile{}, err
	}

	if exists {
		return model.Profile{}, model.ErrConflict
	}

	query = `INSERT INTO dueit.m_profiles (id, user_id, quotes, created_at, created_by, updated_at) 
			 VALUES ($1, $2, $3, $4, $5, $6)`

	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.Profile{}, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if _, err := execSTMT.ExecContext(
		ctx,
		profile.ProfileID,
		profile.UserID,
		profile.Quote,
		profile.CreatedAt,
		profile.CreatedBy,
		profile.UpdatedAt,
	); err != nil {
		log.Err(err).Msg(message.ErrExecDB)
		return model.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepoImpl) UpdateProfile(ctx context.Context, profile model.Profile) (model.Profile, error) {
	query := `UPDATE dueit.m_profiles SET quotes = $1, updated_by = $2, updated_at = $3 
              WHERE user_id = $4 AND id = $5 AND deleted_at IS NULL`

	tx, err := repo.GetTx()
	if err != nil {
		return model.Profile{}, err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.Profile{}, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		profile.Quote,
		profile.ProfileID,
		profile.UpdatedAt,
		profile.UserID,
		profile.ProfileID,
	); err != nil {
		log.Err(err).Msg(message.ErrExecDB)
		return model.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepoImpl) scanRow(row *sql.Row) (model.Profile, error) {
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
		log.Err(err).Msg(message.ErrScanRowDB)
		return model.Profile{}, err
	}
	return profile, nil
}
