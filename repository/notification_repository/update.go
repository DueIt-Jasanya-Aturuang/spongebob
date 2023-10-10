package notification_repository

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (n *NotificationRepositoryImpl) Update(ctx context.Context, req *repository.Notification) error {
	query := `UPDATE dueit.m_notification SET status = $1, updated_at = $2, updated_by = $3
			  WHERE id = $4 and profile_id = $5 AND deleted_at IS NULL`

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
		req.Status,
		req.UpdatedAt,
		req.UpdatedBy,
		req.ID,
		req.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
