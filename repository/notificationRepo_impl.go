package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type NotificationRepoImpl struct {
	domain.UnitOfWorkRepository
}

func NewNotificationRepoImpl(uow domain.UnitOfWorkRepository) domain.NotificationRepo {
	return &NotificationRepoImpl{
		UnitOfWorkRepository: uow,
	}
}

func (n *NotificationRepoImpl) Create(ctx context.Context, req *domain.Notification) error {
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

func (n *NotificationRepoImpl) Update(ctx context.Context, req *domain.Notification) error {
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

func (n *NotificationRepoImpl) Delete(ctx context.Context, id string, profileID string) error {
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

func (n *NotificationRepoImpl) DeleteAllByProfileID(ctx context.Context, profileID string) error {
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

func (n *NotificationRepoImpl) GetAllByProfileID(ctx context.Context, req *domain.RequestGetAllByPaginate) (*[]domain.Notification, error) {
	query := `SELECT id, profile_id, user_config_id, message, title, icon, status, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.m_notification WHERE profile_id=$1 AND deleted_at IS NULL `

	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 10`

	conn, err := n.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	var rows *sql.Rows
	if req.ID != "" {
		rows, err = stmt.QueryContext(ctx, req.ProfileID, req.ID)
	} else {
		rows, err = stmt.QueryContext(ctx, req.ProfileID)
	}

	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, errClose)
		}
	}()

	var notifications []domain.Notification
	var notification domain.Notification

	for rows.Next() {
		if err = rows.Scan(
			&notification.ID,
			&notification.ProfileID,
			&notification.UserConfigID,
			&notification.Message,
			&notification.Title,
			&notification.Icon,
			&notification.Status,
			&notification.CreatedAt,
			&notification.CreatedBy,
			&notification.UpdatedAt,
			&notification.UpdatedBy,
			&notification.DeletedAt,
			&notification.DeletedBy,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	return &notifications, nil
}

func (n *NotificationRepoImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.Notification, error) {
	query := `SELECT id, profile_id, user_config_id, message, title, icon, status, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.m_notification WHERE id=$1 AND profile_id=$2 AND deleted_at IS NULL`

	conn, err := n.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	var notification domain.Notification

	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&notification.ID,
		&notification.ProfileID,
		&notification.UserConfigID,
		&notification.Message,
		&notification.Title,
		&notification.Icon,
		&notification.Status,
		&notification.CreatedAt,
		&notification.CreatedBy,
		&notification.UpdatedAt,
		&notification.UpdatedBy,
		&notification.DeletedAt,
		&notification.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &notification, nil
}

func (n *NotificationRepoImpl) GetNotifHelperByName(ctx context.Context, name string) (*domain.NotificationHelper, error) {
	query := `SELECT id, name, status, title, icon, message, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.h_notify_message WHERE name=$1 AND deleted_at IS NULL LIMIT 1`

	conn, err := n.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	var notificationHelper domain.NotificationHelper

	if err = stmt.QueryRowContext(ctx, name).Scan(
		&notificationHelper.ID,
		&notificationHelper.Name,
		&notificationHelper.Status,
		&notificationHelper.Title,
		&notificationHelper.Icon,
		&notificationHelper.Message,
		&notificationHelper.CreatedAt,
		&notificationHelper.CreatedBy,
		&notificationHelper.UpdatedAt,
		&notificationHelper.UpdatedBy,
		&notificationHelper.DeletedAt,
		&notificationHelper.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &notificationHelper, nil
}
