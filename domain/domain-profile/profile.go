package domainprofile

import (
	"context"
	"database/sql"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileById(context.Context, *sql.DB, string) (*Profile, error)
	GetProfileByUserId(context.Context, *sql.DB, string) (*Profile, error)
	StoreProfile(context.Context, *sql.Tx, Profile) (*Profile, error)
	UpdateProfile(context.Context, *sql.Tx, Profile) (*Profile, error)
}

type ProfileUsecase interface {
	GetProfileById(context.Context, string) (*ProfileResp, error)
	StoreProfile(context.Context, string) (*Profile, error)
}
