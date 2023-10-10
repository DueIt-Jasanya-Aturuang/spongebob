package usecase

import (
	"context"
)

type NotificationUsecase interface {
	UpdateStatus(ctx context.Context, id, profileID string) (*ResponseNotification, error)
	DeleteByIDAndProfileID(ctx context.Context, id string, profileID string) error
	DeleteAllByProfileID(ctx context.Context, profileID string) error
	GetAllByProfileID(ctx context.Context, req *RequestGetAllNotification) (*[]ResponseNotification, string, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseNotification, error)
}

type ResponseNotification struct {
	ID           string
	ProfileID    string
	UserConfigID string
	Message      string
	Title        string
	Icon         string
	Status       string
	CreatedAt    int64
}

type RequestGetAllNotification struct {
	ID        string
	ProfileID string
	Order     string
}
