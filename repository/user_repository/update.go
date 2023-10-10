package user_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (u *UserRepoImpl) Update(ctx context.Context, user *repository.User) error {
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

func (u *UserRepoImpl) UpdateUsername(ctx context.Context, user *repository.User) (bool, error) {
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
