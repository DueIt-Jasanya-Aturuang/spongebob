package notification_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type NotificationRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewNotificationRepositoryImpl(uow repository.UnitOfWorkRepository) repository.NotificationRepository {
	return &NotificationRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
