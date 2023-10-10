package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	CheckPhoneNumberExists(ctx context.Context, id string, newPhoneNumber string) (bool, error)
	UpdateUsername(ctx context.Context, user *User) (bool, error)
	UnitOfWorkRepository
}

type User struct {
	ID              string
	FullName        string
	Gender          string
	Image           string
	Username        string
	Email           string
	Password        string
	PhoneNumber     sql.NullString
	EmailVerifiedAt bool
	AuditInfo
}
