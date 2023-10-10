package profileConfig_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileConfigRepositoryImpl) Create(ctx context.Context, profileCfg *repository.ProfileConfig) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM dueit.m_user_config WHERE profile_id = $1 AND config_name = $2)`
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

	if err = querySTMT.QueryRowContext(ctx, profileCfg.ProfileID, profileCfg.ConfigName).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	if exist {
		return true, nil
	}

	query = `INSERT INTO dueit.m_user_config 
    					 (id, profile_id, config_name, config_value, status, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

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

	_, err = execSTMT.ExecContext(
		ctx,
		profileCfg.ID,
		profileCfg.ProfileID,
		profileCfg.ConfigName,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.CreatedAt,
		profileCfg.CreatedBy,
		profileCfg.UpdatedAt,
	)

	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
	}

	return false, err
}
