package domain

import (
	"context"
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
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
	GetByID(ctx context.Context, id string) (Profile, error)
	GetByUserID(ctx context.Context, userID string) (Profile, error)
	Store(ctx context.Context, profile Profile) error
	Update(ctx context.Context, profile Profile) error
	UnitOfWorkRepository
}

func (p *Profile) DefaultValue(userID string) *Profile {
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

func (p *Profile) ToResp() *ResponseProfile {
	var quote string
	var profesi string
	if p.Quote.Valid {
		quote = p.Quote.String
	} else {
		quote = "null"
	}

	if p.Profesi.Valid {
		profesi = p.Profesi.String
	} else {
		profesi = "null"
	}

	return &ResponseProfile{
		ProfileID: p.ProfileID,
		Quote:     quote,
		Profesi:   profesi,
	}
}
