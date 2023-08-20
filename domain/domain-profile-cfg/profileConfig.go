package domainprofilecfg

import (
	"context"
	"database/sql"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileCfgRepo
type ProfileCfgRepo interface {
	StoreProfileCfg(context.Context, *sql.Tx, ProfileCfg) error
	UpdateProfileCfg(context.Context, *sql.Tx, ProfileCfg) error
	GetProfileCfgById(context.Context, *sql.DB, string) (*ProfileCfg, error)
	GetProfileCfgByScheduler(context.Context, *sql.DB, ProfileCfgScheduler) (*[]ProfileCfg, error)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileCfgUsecase
type ProfileCfgUsecase interface {
	CreateProfileCfg(context.Context, ProfileCfgReq) (*ProfileCfgResp, error)
	GetProfileCfgById(context.Context, string) (*ProfileCfgResp, error)
	UpdateProfileCfg(context.Context, ProfileCfgReq, string) (*ProfileCfgResp, error)
}
