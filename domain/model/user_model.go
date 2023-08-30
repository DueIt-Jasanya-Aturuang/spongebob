package model

import (
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

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
	CreatedAt       int64
	CreatedBy       string
	UpdatedAt       int64
	UpdatedBy       sql.NullString
	DeletedAt       sql.NullInt64
	DeletedBy       sql.NullString
}

func (u *User) ToResp(emailFormat string) *dto.UserResp {
	var phoneNumber string

	if u.PhoneNumber.Valid {
		phoneNumber = u.PhoneNumber.String
	} else {
		phoneNumber = "null"
	}

	return &dto.UserResp{
		ID:              u.ID,
		FullName:        u.FullName,
		Gender:          u.Gender,
		Image:           u.Image,
		Username:        u.Username,
		Email:           u.Email,
		EmailFormat:     emailFormat,
		PhoneNumber:     phoneNumber,
		EmailVerifiedAt: u.EmailVerifiedAt,
	}
}
