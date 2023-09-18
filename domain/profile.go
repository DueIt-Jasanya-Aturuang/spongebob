package domain

import (
	"context"
	"database/sql"
)

type Profile struct {
	ProfileID string
	UserID    string
	Quote     sql.NullString
	Profesi   sql.NullString
	AuditInfo
}

//counterfeiter:generate -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetByID(ctx context.Context, id string) (*Profile, error)
	GetByUserID(ctx context.Context, userID string) (*Profile, error)
	Create(ctx context.Context, profile *Profile) (bool, error)
	Update(ctx context.Context, profile *Profile) error
	UnitOfWorkRepository
}
