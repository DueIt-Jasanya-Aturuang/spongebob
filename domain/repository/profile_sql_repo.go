package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileById(context.Context, string) (*model.Profile, error)
	GetProfileByUserId(context.Context, string) (*model.Profile, error)
	StoreProfile(context.Context, *sql.Tx, model.Profile) (model.Profile, error)
	UpdateProfile(context.Context, *sql.Tx, model.Profile) (*model.Profile, error)
}
