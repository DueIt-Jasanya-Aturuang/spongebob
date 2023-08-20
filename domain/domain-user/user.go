package domainuser

import (
	"context"
	"database/sql"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . UserRepo
type UserRepo interface {
	GetUserById(context.Context, *sql.DB, string) (*User, error)
	UpdateUser(context.Context, *sql.Tx, User) (*User, error)
	UpdateUsername(context.Context, *sql.Tx, User) (*User, error)
}
