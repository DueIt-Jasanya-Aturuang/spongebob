package profile_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileRepositoryImpl) Update(ctx context.Context, profile *repository.Profile) error {
	query := `UPDATE dueit.m_profiles SET quotes = $1, profesi = $2,  updated_by = $3, updated_at = $4 
              WHERE user_id = $5 AND id = $6 AND deleted_at IS NULL`

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

	if _, err = stmt.ExecContext(
		ctx,
		profile.Quote,
		profile.Profesi,
		profile.ProfileID,
		profile.UpdatedAt,
		profile.UserID,
		profile.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
