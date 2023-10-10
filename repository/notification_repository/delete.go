package notification_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (n *NotificationRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	query := `UPDATE dueit.m_notification SET deleted_at = $1, deleted_by = $2
			  WHERE id = $3 and profile_id = $4 AND deleted_at IS NULL`

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
		time.Now().Unix(),
		profileID,
		id,
		profileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (n *NotificationRepositoryImpl) DeleteAllByProfileID(ctx context.Context, profileID string) error {
	query := `UPDATE dueit.m_notification SET deleted_at = $1, deleted_by = $2
			  WHERE profile_id = $3 AND deleted_at IS NULL`

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
		time.Now().Unix(),
		profileID,
		profileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}
