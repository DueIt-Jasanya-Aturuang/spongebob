package repositories

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/rs/zerolog/log"
)

type UserRepoImpl struct{}

func NewUserRepoImpl() domain.UserRepo {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) scanRow(row *sql.Row) (*domain.User, error) {
	var user domain.User
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
		log.Err(err).Msg(domain.LogErrScanning)
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepoImpl) GetUserById(ctx context.Context, db *sql.DB, id string) (*domain.User, error) {
	query := "SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM auth.m_users WHERE id = $1 AND deleted_at IS $2"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domain.LogErrSTMT)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id, "NULL")

	user, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepoImpl) UpdateUser(ctx context.Context, tx *sql.Tx, entity domain.User) (*domain.User, error) {
	query := "UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, updated_at = $5, updated_by = $6 WHERE id = $7 AND deleted_at IS NULL"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domain.LogErrSTMT)
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
		log.Err(err).Msg(domain.LogErrExec)
		return nil, err
	}

	return &entity, nil
}

func (repo *UserRepoImpl) UpdateUsername(ctx context.Context, tx *sql.Tx, entity domain.User) (*domain.User, error) {
	query := "SELECT username, id FROM auth.m_users WHERE username=$1 AND id<>$2 AND deleted_at IS NULL"
	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domain.LogErrSTMT)
		return nil, err
	}

	rows, err := querySTMT.QueryContext(ctx, entity.Username, entity.ID)
	if err != nil {
		log.Err(err).Msg(domain.LogErrQuery)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		return nil, domain.ErrUsernameAlvailable
	}

	query = "UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 WHERE id = $4 AND deleted_at IS NULL"
	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Err(err).Msg(domain.LogErrSTMT)
		return nil, err
	}

	if _, err := execSTMT.ExecContext(
		ctx,
		entity.Username,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.ID,
	); err != nil {
		log.Err(err).Msg(domain.LogErrExec)
		return nil, err
	}

	return &entity, nil
}
