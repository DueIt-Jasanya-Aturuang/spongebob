package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . UserRepo
type UserRepo interface {
	GetUserByID(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, model.User) (*model.User, error)
	UpdateUsername(context.Context, model.User) (*model.User, error)
	UoW() UnitOfWork
}
