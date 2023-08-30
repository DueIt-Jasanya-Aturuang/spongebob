package repository

import (
	"context"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileByID(c context.Context, id string) (model.Profile, error)
	GetProfileByUserID(c context.Context, userID string) (model.Profile, error)
	StoreProfile(c context.Context, profile model.Profile) (model.Profile, error)
	// UpdateProfile update profile, setup query update
	// get tx from start transaction in usecase layer
	// start prepared statement for query update profile
	// defer close prepared statement
	// exec context for update profile
	UpdateProfile(c context.Context, profile model.Profile) (model.Profile, error)
	UnitOfWork
}
