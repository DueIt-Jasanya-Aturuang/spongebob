package notification_usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (n *NotificationUsecaseImpl) UpdateStatus(ctx context.Context, id, profileID string) (*usecase.ResponseNotification, error) {
	notification, err := n.notifRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.NotificationNotFound
		}
		return nil, err
	}

	notification.Status = "read"
	notification.UpdatedAt = time.Now().Unix()
	notification.UpdatedBy = repository.NewNullString(profileID)

	err = n.notifRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = n.notifRepo.Update(ctx, notification)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponseNotification{
		ID:           notification.ID,
		ProfileID:    notification.ProfileID,
		UserConfigID: notification.UserConfigID,
		Message:      notification.Message,
		Title:        notification.Title,
		Icon:         notification.Icon,
		Status:       notification.Status,
		CreatedAt:    notification.CreatedAt,
	}
	return resp, nil
}
