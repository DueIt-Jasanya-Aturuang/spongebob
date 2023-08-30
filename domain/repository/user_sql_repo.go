package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . UserRepo
type UserRepo interface {
	GetUserByID(c context.Context, id string) (model.User, error)
	UpdateUser(c context.Context, user model.User) (model.User, error)
	CheckPhoneNumberExists(c context.Context, id string, newPhoneNumber string) (exists bool, err error)
	UpdateUsername(c context.Context, user model.User) (model.User, error)
	UnitOfWork
}
