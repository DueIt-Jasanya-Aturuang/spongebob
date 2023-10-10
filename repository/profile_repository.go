package repository

import (
	"context"
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

//counterfeiter:generate -o ./../mocks . ProfileRepo
type ProfileRepository interface {
	GetByID(ctx context.Context, id string) (*Profile, error)
	GetByUserID(ctx context.Context, userID string) (*Profile, error)
	Create(ctx context.Context, profile *Profile) (bool, error)
	Update(ctx context.Context, profile *Profile) error
	UnitOfWorkRepository
}

type Profile struct {
	ProfileID string
	UserID    string
	Quote     sql.NullString
	Profesi   sql.NullString
	AuditInfo
}

func DefaultValueProfile(userID string) *Profile {
	id := uuid.NewV4().String()
	return &Profile{
		ProfileID: id,
		UserID:    userID,
		Quote:     sql.NullString{},
		Profesi:   sql.NullString{},
		AuditInfo: AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: id,
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}
}
