package domainprofile

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Profile struct {
	ProfileId string
	UserId    string
	Quote     sql.NullString
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}

func (p *Profile) DefaultValue() Profile {
	id := uuid.NewV4().String()
	return Profile{
		ProfileId: id,
		UserId:    p.UserId,
		Quote:     sql.NullString{},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}
}

func (p *Profile) ToResp() ProfileResp {
	return ProfileResp{
		ProfileID: p.ProfileId,
		Quote:     p.Quote,
	}
}
