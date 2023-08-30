package repository

import (
	"context"
	"database/sql"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"github.com/rs/zerolog/log"
)

type UserRepoImpl struct {
	repository.UnitOfWork
}

func NewUserRepoImpl(uow repository.UnitOfWork) repository.UserRepo {
	return &UserRepoImpl{
		UnitOfWork: uow,
	}
}

func (repo *UserRepoImpl) GetUserByID(
	ctx context.Context, id string,
) (model.User, error) {
	query := `SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, 
       		  		 created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
			  FROM auth.m_users WHERE id = $1 AND deleted_at IS NULL`

	conn, err := repo.GetConn()
	if err != nil {
		return model.User{}, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.User{}, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	row := stmt.QueryRowContext(ctx, id)

	user, err := repo.scanRow(row)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepoImpl) CheckPhoneNumberExists(
	c context.Context, id string, newPhoneNumber string,
) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM auth.m_users WHERE phone_number = $1 AND id<>$2 AND deleted_at IS NULL)`
	// get connection database from uow
	conn, err := repo.GetConn()
	if err != nil {
		return false, err
	}

	stmt, err := conn.PrepareContext(c, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return false, err
	}
	defer func() {
		if errStmt := stmt.Close(); errStmt != nil {
			log.Err(errStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if err = stmt.QueryRowContext(c, newPhoneNumber, id).Scan(&exists); err != nil {
		log.Err(err).Msg(message.ErrQueryRowDB)
		return false, err
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (repo *UserRepoImpl) UpdateUser(
	ctx context.Context, entity model.User,
) (model.User, error) {
	query := `UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, 
                        			  updated_at = $5, updated_by = $6 
              WHERE id = $7 AND deleted_at IS NULL`

	tx, err := repo.GetTx()
	if err != nil {
		return model.User{}, err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.User{}, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if _, err := stmt.ExecContext(
		ctx,
		entity.FullName,
		entity.Gender,
		entity.Image,
		entity.PhoneNumber,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(message.ErrExecDB)
		return model.User{}, err
	}

	return entity, nil
}

func (repo *UserRepoImpl) UpdateUsername(
	ctx context.Context, entity model.User,
) (model.User, error) {
	query := `SELECT EXISTS (SELECT 1 FROM auth.m_users WHERE username = $1 AND id<>$2 AND deleted_at IS NULL)`
	var exists bool

	tx, err := repo.GetTx()
	if err != nil {
		return model.User{}, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.User{}, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Err(errQueryStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if err = querySTMT.QueryRowContext(ctx, entity.Username, entity.ID).Scan(&exists); err != nil {
		log.Err(err).Msg(message.ErrQueryRowDB)
		return model.User{}, err
	}

	if exists {
		return model.User{}, model.ErrConflict
	}

	query = `UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 
             WHERE id = $4 AND deleted_at IS NULL`

	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(message.ErrOpenStmtDB)
		return model.User{}, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Err(errExecStmt).Msg(message.ErrCloseStmtDB)
		}
	}()

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.Username,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(message.ErrExecDB)
		return model.User{}, err
	}

	return entity, nil
}

func (repo *UserRepoImpl) scanRow(row *sql.Row) (model.User, error) {
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
		log.Err(err).Msg(message.ErrScanRowDB)
		return model.User{}, err
	}
	return user, nil
}
