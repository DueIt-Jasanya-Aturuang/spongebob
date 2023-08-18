package domain

import (
	"context"
	"database/sql"
)

// user resp
type UserResp struct {
	ID              string  `json:"id"`
	FullName        string  `json:"full_name"`
	Gender          string  `json:"gender"`
	Image           string  `json:"image"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	EmailFormat     string  `json:"email_format"`
	PhoneNumber     *string `json:"phone_number"`
	EmailVerifiedAt bool    `json:"activited"`
}

// user entities
type User struct {
	ID              string         `json:"user_id"`
	FullName        string         `json:"full_name"`
	Gender          string         `json:"gender"`
	Image           string         `json:"image"`
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	PhoneNumber     sql.NullString `json:"phone_number"`
	EmailVerifiedAt bool           `json:"email_verified_at"`
	CreatedAt       int64          `json:"created_at"`
	CreatedBy       string         `json:"created_by"`
	UpdatedAt       int64          `json:"updated_at"`
	UpdatedBy       sql.NullString `json:"updated_by"`
	DeletedAt       sql.NullInt64  `json:"deleted_at"`
	DeletedBy       sql.NullString `json:"deleted_by"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./mocks . UserRepo
type UserRepo interface {
	GetUserById(context.Context, *sql.DB, string) (*User, error)
	UpdateUser(context.Context, *sql.Tx, User) (*User, error)
	UpdateUsername(context.Context, *sql.Tx, User) (*User, error)
}
