package notification_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

func (n *NotificationUsecaseImpl) DeleteByIDAndProfileID(ctx context.Context, id string, profileID string) error {
	err := n.notifRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err := n.notifRepo.Delete(ctx, id, profileID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (n *NotificationUsecaseImpl) DeleteAllByProfileID(ctx context.Context, profileID string) error {
	err := n.notifRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err := n.notifRepo.DeleteAllByProfileID(ctx, profileID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
