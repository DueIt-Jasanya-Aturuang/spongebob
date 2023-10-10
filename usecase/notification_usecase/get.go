package notification_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

func (n *NotificationUsecaseImpl) GetAllByProfileID(ctx context.Context, req *usecase.RequestGetAllNotification) (*[]usecase.ResponseNotification, string, error) {
	order, operation := usecase.GetOrder(req.Order)
	notifications, err := n.notifRepo.GetAllByProfileID(ctx, &repository.RequestGetAllNotification{
		ProfileID: req.ProfileID,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*notifications) < 1 {
		return nil, "", nil
	}

	var responses []usecase.ResponseNotification
	var response *usecase.ResponseNotification

	for _, notification := range *notifications {
		response = &usecase.ResponseNotification{
			ID:           notification.ID,
			ProfileID:    notification.ProfileID,
			UserConfigID: notification.UserConfigID,
			Message:      notification.Message,
			Title:        notification.Title,
			Icon:         notification.Icon,
			Status:       notification.Status,
			CreatedAt:    notification.CreatedAt,
		}
		responses = append(responses, *response)
	}

	cursor := (*notifications)[len(*notifications)-1].ID
	return &responses, cursor, nil
}

func (n *NotificationUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*usecase.ResponseNotification, error) {
	notification, err := n.notifRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.NotificationNotFound
		}
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
