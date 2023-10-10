package repository

import (
	"context"
)

type NotificationRepository interface {
	Create(ctx context.Context, req *Notification) error
	Update(ctx context.Context, req *Notification) error
	Delete(ctx context.Context, id string, profileID string) error
	DeleteAllByProfileID(ctx context.Context, profileID string) error
	GetAllByProfileID(ctx context.Context, req *RequestGetAllNotification) (*[]Notification, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*Notification, error)
	GetNotifHelperByName(ctx context.Context, name string) (*NotificationHelper, error)
	UnitOfWorkRepository
}

type RequestGetAllNotification struct {
	ProfileID string
	InfiniteScrollData
}

type Notification struct {
	ID           string
	ProfileID    string
	UserConfigID string
	Message      string
	Title        string
	Icon         string
	Status       string
	AuditInfo
}

type NotificationHelper struct {
	ID      string
	Name    string
	Status  string
	Title   string
	Icon    string
	Message string
	AuditInfo
}
