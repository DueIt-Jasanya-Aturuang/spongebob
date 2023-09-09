package model

import (
	"database/sql"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	uuid "github.com/satori/go.uuid"
)

type Profile struct {
	ProfileID string
	UserID    string
	Quote     sql.NullString
	Profesi   sql.NullString
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}

func (p *Profile) DefaultValue(userID string) *Profile {
	id := uuid.NewV4().String()
	return &Profile{
		ProfileID: id,
		UserID:    userID,
		Quote:     sql.NullString{},
		Profesi:   sql.NullString{},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}
}

func (p *Profile) ToResp() *dto.ProfileResp {
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

	return &dto.ProfileResp{
		ProfileID: p.ProfileID,
		Quote:     quote,
		Profesi:   profesi,
	}
}
