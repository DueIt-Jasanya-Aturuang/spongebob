package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileByID(context.Context, string) (*model.Profile, error)
	GetProfileByUserID(context.Context, string) (*model.Profile, error)
	StoreProfile(context.Context, model.Profile) (model.Profile, error)
	UpdateProfile(context.Context, model.Profile) (*model.Profile, error)
	UoW() UnitOfWork
}
