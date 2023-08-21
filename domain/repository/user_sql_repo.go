package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . UserRepo
type UserRepo interface {
	GetUserById(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, *sql.Tx, model.User) (*model.User, error)
	UpdateUsername(context.Context, *sql.Tx, model.User) (*model.User, error)
}
