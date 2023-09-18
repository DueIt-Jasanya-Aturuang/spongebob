package domain

import (
	"context"
	"database/sql"
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
	AuditInfo
}

func (u *User) ToResp(emailFormat string) *ResponseUser {
	var phoneNumber string

	if u.PhoneNumber.Valid {
		phoneNumber = u.PhoneNumber.String
	} else {
		phoneNumber = "null"
	}

	return &ResponseUser{
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

type UserRepo interface {
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, user User) (User, error)
	CheckPhoneNumberExists(c context.Context, id string, newPhoneNumber string) (exists bool, err error)
	UpdateUsername(c context.Context, user User) (User, error)
	UnitOfWorkRepository
}
