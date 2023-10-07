package domain

import (
	"context"
)

type NotificationRepo interface {
	Create(ctx context.Context, req *Notification) error
	Update(ctx context.Context, req *Notification) error
	Delete(ctx context.Context, id string, profileID string) error
	DeleteAllByProfileID(ctx context.Context, profileID string) error
	GetAllByProfileID(ctx context.Context, req *RequestGetAllByPaginate) (*[]Notification, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*Notification, error)
	GetNotifHelperByName(ctx context.Context, name string) (*NotificationHelper, error)
	UnitOfWorkRepository
}

type NotificationUsecase interface {
	UpdateStatus(ctx context.Context, id, profileID string) (*ResponseNotification, error)
	DeleteByIDAndProfileID(ctx context.Context, id string, profileID string) error
	DeleteAllByProfileID(ctx context.Context, profileID string) error
	GetAllByProfileID(ctx context.Context, req *RequestGetAllByPaginate) (*[]ResponseNotification, string, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseNotification, error)
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

type ResponseNotification struct {
	ID           string `json:"id"`
	ProfileID    string `json:"profile_id"`
	UserConfigID string `json:"user_config_id"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Icon         string `json:"icon"`
	Status       string `json:"status"`
	CreatedAt    int64  `json:"created_at"`
}
