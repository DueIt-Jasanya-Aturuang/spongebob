package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/helpers"
)

type NotificationUsecaseImpl struct {
	notifRepo domain.NotificationRepo
}

func NewNotificationUsecaseImpl(
	notifRepo domain.NotificationRepo,
) domain.NotificationUsecase {
	return &NotificationUsecaseImpl{
		notifRepo: notifRepo,
	}
}

func (n *NotificationUsecaseImpl) UpdateStatus(ctx context.Context, id, profileID string) (*domain.ResponseNotification, error) {
	err := n.notifRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer n.notifRepo.CloseConn()

	notif, err := n.notifRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotificationNotFound
		}
		return nil, err
	}

	notif.Status = "read"
	notif.UpdatedAt = time.Now().Unix()
	notif.UpdatedBy = helpers.NewNullString(profileID)

	err = n.notifRepo.Update(ctx, notif)
	if err != nil {
		return nil, err
	}

	resp := converter.NotifModelToResponse(notif)
	return resp, nil
}

func (n *NotificationUsecaseImpl) DeleteByIDAndProfileID(ctx context.Context, id string, profileID string) error {
	err := n.notifRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	defer n.notifRepo.CloseConn()

	err = n.notifRepo.Delete(ctx, id, profileID)
	if err != nil {
		return err
	}

	return nil
}

func (n *NotificationUsecaseImpl) DeleteAllByProfileID(ctx context.Context, profileID string) error {
	err := n.notifRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	defer n.notifRepo.CloseConn()

	err = n.notifRepo.DeleteAllByProfileID(ctx, profileID)
	if err != nil {
		return err
	}

	return nil
}

func (n *NotificationUsecaseImpl) GetAllByProfileID(ctx context.Context, req *domain.RequestGetAllByPaginate) (*[]domain.ResponseNotification, string, error) {
	err := n.notifRepo.OpenConn(ctx)
	if err != nil {
		return nil, "", err
	}
	defer n.notifRepo.CloseConn()

	notifications, err := n.notifRepo.GetAllByProfileID(ctx, req)
	if err != nil {
		return nil, "", err
	}

	var responses []domain.ResponseNotification
	var response domain.ResponseNotification

	for _, notification := range *notifications {
		response = *converter.NotifModelToResponse(&notification)
		responses = append(responses, response)
	}

	cursor := (*notifications)[len(*notifications)-1].ID
	return &responses, cursor, nil
}

func (n *NotificationUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseNotification, error) {
	err := n.notifRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer n.notifRepo.CloseConn()

	notif, err := n.notifRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotificationNotFound
		}
		return nil, err
	}

	resp := converter.NotifModelToResponse(notif)

	return resp, nil

}
