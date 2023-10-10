package notification_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (n *NotificationRepositoryImpl) GetAllByProfileID(ctx context.Context, req *repository.RequestGetAllNotification) (*[]repository.Notification, error) {
	query := `SELECT id, profile_id, user_config_id, message, title, icon, status, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.m_notification WHERE profile_id=$1 AND deleted_at IS NULL `

	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 10`

	db, err := n.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
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

	var notifications []repository.Notification
	var notification repository.Notification

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

func (n *NotificationRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*repository.Notification, error) {
	query := `SELECT id, profile_id, user_config_id, message, title, icon, status, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.m_notification WHERE id=$1 AND profile_id=$2 AND deleted_at IS NULL`

	db, err := n.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	var notification repository.Notification

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

func (n *NotificationRepositoryImpl) GetNotifHelperByName(ctx context.Context, name string) (*repository.NotificationHelper, error) {
	query := `SELECT id, name, status, title, icon, message, created_at, created_by, updated_at,
       					updated_by, deleted_at, deleted_by
				FROM dueit.h_notify_message WHERE name=$1 AND deleted_at IS NULL LIMIT 1`

	db, err := n.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errExecStmt := stmt.Close(); errExecStmt != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errExecStmt)
		}
	}()

	var notificationHelper repository.NotificationHelper

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
