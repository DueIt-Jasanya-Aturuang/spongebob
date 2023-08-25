package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileByID(c context.Context, id string) (*model.Profile, error)
	GetProfileByUserID(c context.Context, userID string) (*model.Profile, error)
	StoreProfile(c context.Context, profile model.Profile) (model.Profile, error)
	UpdateProfile(c context.Context, profile model.Profile) (*model.Profile, error)
	UoW() UnitOfWork
}
