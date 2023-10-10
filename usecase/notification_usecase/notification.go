package notification_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

type NotificationUsecaseImpl struct {
	notifRepo repository.NotificationRepository
}

func NewNotificationUsecaseImpl(
	notifRepo repository.NotificationRepository,
) usecase.NotificationUsecase {
	return &NotificationUsecaseImpl{
		notifRepo: notifRepo,
	}
}
