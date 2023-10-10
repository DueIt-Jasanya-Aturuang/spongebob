package profile_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileRepositoryImpl) Create(ctx context.Context, profile *repository.Profile) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_profiles WHERE user_id = $1 AND deleted_at IS NULL)`
	var exist bool

	tx, err := p.GetTx()
	if err != nil {
		return false, err
	}

	querySTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errQueryStmt := querySTMT.Close(); errQueryStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errQueryStmt)
		}
	}()

	if err = querySTMT.QueryRowContext(ctx, profile.UserID).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	query = `INSERT INTO dueit.m_profiles (id, user_id, quotes, profesi, created_at, created_by, updated_at) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	execSTMT, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errExecStmt := execSTMT.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	if _, err = execSTMT.ExecContext(
		ctx,
		profile.ProfileID,
		profile.UserID,
		profile.Quote,
		profile.Profesi,
		profile.CreatedAt,
		profile.CreatedBy,
		profile.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return false, err
	}

	return false, nil
}
