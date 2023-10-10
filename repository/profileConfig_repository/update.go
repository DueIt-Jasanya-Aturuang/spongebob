package profileConfig_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileConfigRepositoryImpl) Update(ctx context.Context, profileCfg *repository.ProfileConfig) error {
	query := `UPDATE dueit.m_user_config SET config_value = $1, status = $2, updated_at = $3, updated_by = $4
			  WHERE id = $5 and profile_id = $6`

	tx, err := p.GetTx()
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

	_, err = stmt.ExecContext(
		ctx,
		profileCfg.ConfigValue,
		profileCfg.Status,
		profileCfg.UpdatedAt,
		profileCfg.UpdatedBy,
		profileCfg.ID,
		profileCfg.ProfileID,
	)

	if err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
	}

	return err
}
