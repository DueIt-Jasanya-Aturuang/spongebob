package repositories

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exceptions"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/rs/zerolog/log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepoImpl(db *sql.DB) repository.UserRepo {
	return &UserRepoImpl{
		db: db,
	}
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
		log.Err(err).Msg(exceptions.LogErrScanning)
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepoImpl) GetUserById(ctx context.Context, id string) (*model.User, error) {
	query := "SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM auth.m_users WHERE id = $1 AND deleted_at IS NULL"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)

	user, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepoImpl) UpdateUser(ctx context.Context, tx *sql.Tx, entity model.User) (*model.User, error) {
	query := "SELECT phone_number, id FROM auth.m_users WHERE phone_number=$1 AND id<>$2 AND deleted_at IS NULL"
	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	rows, err := querySTMT.QueryContext(ctx, entity.PhoneNumber, entity.ID)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrQuery)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		log.Info().Msg(exceptions.ErrPhoneAlvailable.Error())
		return nil, exceptions.ErrPhoneAlvailable
	}

	query = "UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, updated_at = $5, updated_by = $6 WHERE id = $7 AND deleted_at IS NULL"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

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
		log.Err(err).Msg(exceptions.LogErrExec)
		return nil, err
	}

	return &entity, nil
}

func (repo *UserRepoImpl) UpdateUsername(ctx context.Context, tx *sql.Tx, entity model.User) (*model.User, error) {
	query := "SELECT username, id FROM auth.m_users WHERE username=$1 AND id<>$2 AND deleted_at IS NULL"
	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	rows, err := querySTMT.QueryContext(ctx, entity.Username, entity.ID)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrQuery)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		log.Info().Msg(exceptions.ErrUsernameAlvailable.Error())
		return nil, exceptions.ErrUsernameAlvailable
	}

	query = "UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 WHERE id = $4 AND deleted_at IS NULL"
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(exceptions.LogErrSTMT)
		return nil, err
	}

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.Username,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(exceptions.LogErrExec)
		return nil, err
	}

	return &entity, nil
}
