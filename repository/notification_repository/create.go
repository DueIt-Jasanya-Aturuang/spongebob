package notification_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (n *NotificationRepositoryImpl) Create(ctx context.Context, req *repository.Notification) error {
	query := `INSERT INTO dueit.m_notification 
    					 (id, profile_id, user_config_id, message, title, icon, status, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	tx, err := n.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		req.ID,
		req.ProfileID,
		req.UserConfigID,
		req.Message,
		req.Title,
		req.Icon,
		req.Status,
		req.CreatedAt,
		req.CreatedBy,
		req.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
