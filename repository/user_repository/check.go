package user_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

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
