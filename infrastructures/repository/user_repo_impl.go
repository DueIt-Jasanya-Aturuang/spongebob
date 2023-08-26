package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type UserRepoImpl struct {
	uow repository.UnitOfWork
}

func NewUserRepoImpl(uow repository.UnitOfWork) repository.UserRepo {
	return &UserRepoImpl{
		uow: uow,
	}
}

func (repo *UserRepoImpl) UoW() repository.UnitOfWork {
	return repo.uow
}

func (repo *UserRepoImpl) scanRow(row *sql.Row) (*model.User, error) {
	var user model.User
	if err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Gender,
		&user.Image,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.EmailVerifiedAt,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
		&user.DeletedAt,
		&user.DeletedBy,
	); err != nil {
		log.Err(err).Msg(exception.LogErrDBScanning)
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepoImpl) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	query := "SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, " +
		"created_at, created_by, updated_at, updated_by, deleted_at, deleted_by " +
		"FROM auth.m_users WHERE id = $1 AND deleted_at IS NULL"

	conn, err := repo.uow.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if errConn := conn.Close(); errConn != nil {
			log.Err(errConn).Msg(exception.LogErrDBCloseConn)
		}
	}()

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBStmt)
		return nil, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(exception.LogErrDBCloseStmt)
		}
	}()

	row := stmt.QueryRowContext(ctx, id)

	user, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepoImpl) UpdateUser(ctx context.Context, entity model.User) (*model.User, error) {
	query := "SELECT phone_number, id FROM auth.m_users WHERE phone_number=$1 AND id<>$2 AND deleted_at IS NULL"
	tx, err := repo.uow.GetTx()
	if err != nil {
		return nil, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBStmt)
		return nil, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(exception.LogErrDBCloseStmt)
		}
	}()

	rows, err := querySTMT.QueryContext(ctx, entity.PhoneNumber, entity.ID)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBQuery)
		return nil, err
	}
	defer func() {
		if errRows := rows.Close(); errRows != nil {
			log.Err(errRows).Msg(exception.LogErrDBCloseRows)
		}
	}()

	if rows.Next() {
		log.Info().Msg(exception.Err400PhoneAlvailable.Error())
		return nil, exception.Err400PhoneAlvailable
	}

	query = "UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, updated_at = $5, updated_by = $6 " +
		"WHERE id = $7 AND deleted_at IS NULL"
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBStmt)
		return nil, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(exception.LogErrDBCloseStmt)
		}
	}()

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.FullName,
		entity.Gender,
		entity.Image,
		entity.PhoneNumber,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(exception.LogErrDBExec)
		return nil, err
	}

	return &entity, nil
}

func (repo *UserRepoImpl) UpdateUsername(ctx context.Context, entity model.User) (*model.User, error) {
	query := "SELECT username, id FROM auth.m_users WHERE username=$1 AND id<>$2 AND deleted_at IS NULL"
	tx, err := repo.uow.GetTx()
	if err != nil {
		return nil, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBStmt)
		return nil, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(exception.LogErrDBCloseStmt)
		}
	}()

	rows, err := querySTMT.QueryContext(ctx, entity.Username, entity.ID)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBQuery)
		return nil, err
	}
	defer func() {
		if errRows := rows.Close(); errRows != nil {
			log.Err(errRows).Msg(exception.LogErrDBCloseRows)
		}
	}()

	if rows.Next() {
		log.Info().Msg(exception.Err400UsernameAlvailable.Error())
		return nil, exception.Err400UsernameAlvailable
	}

	query = "UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 WHERE id = $4 AND deleted_at IS NULL"

	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exception.LogErrDBStmt)
		return nil, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(exception.LogErrDBCloseStmt)
		}
	}()

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.Username,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(exception.LogErrDBExec)
		return nil, err
	}

	return &entity, nil
}
