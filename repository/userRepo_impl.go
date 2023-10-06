package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type UserRepoImpl struct {
	domain.UnitOfWorkRepository
}

func NewUserRepoImpl(uow domain.UnitOfWorkRepository) domain.UserRepo {
	return &UserRepoImpl{
		UnitOfWorkRepository: uow,
	}
}

func (u *UserRepoImpl) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, fullname, gender, image, username, email, password, phone_number, email_verified_at, 
       		  		 created_at, created_by, updated_at, updated_by, deleted_at, deleted_by 
			  FROM auth.m_users WHERE id = $1 AND deleted_at IS NULL`

	db, err := u.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var user domain.User
	if err = stmt.QueryRowContext(ctx, id).Scan(
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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
	}

	return &user, nil
}

func (u *UserRepoImpl) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE auth.m_users SET fullname = $1, gender = $2, image = $3, phone_number = $4, 
                        			  updated_at = $5, updated_by = $6 
              WHERE id = $7 AND deleted_at IS NULL`

	tx, err := u.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		user.FullName,
		user.Gender,
		user.Image,
		user.PhoneNumber,
		user.UpdatedAt,
		user.UpdatedBy,
		user.ID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (u *UserRepoImpl) CheckPhoneNumberExists(ctx context.Context, id string, newPhoneNumber string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM auth.m_users WHERE phone_number = $1 AND id<>$2 AND deleted_at IS NULL)`
	var exist bool

	db, err := u.GetDB()
	if err != nil {
		return false, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if err = stmt.QueryRowContext(ctx, newPhoneNumber, id).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	return false, nil
}

func (u *UserRepoImpl) UpdateUsername(ctx context.Context, user *domain.User) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM auth.m_users WHERE username = $1 AND id<>$2 AND deleted_at IS NULL)`
	var exist bool

	tx, err := u.GetTx()
	if err != nil {
		return false, err
	}

	queryStmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := queryStmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if err = queryStmt.QueryRowContext(ctx, user.Username, user.ID).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	query = `UPDATE auth.m_users SET username = $1, updated_at = $2, updated_by = $3 
             WHERE id = $4 AND deleted_at IS NULL`

	execStmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := execStmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = execStmt.ExecContext(
		ctx,
		user.Username,
		user.UpdatedAt,
		user.UpdatedBy,
		user.ID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return false, err
	}

	return false, nil
}
