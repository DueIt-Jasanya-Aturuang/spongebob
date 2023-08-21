package repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileCfgRepo
type ProfileCfgRepo interface {
	StoreProfileCfg(context.Context, *sql.Tx, model.ProfileCfg) error
	UpdateProfileCfg(context.Context, *sql.Tx, model.ProfileCfg) error
	GetProfileCfgById(context.Context, string) (*model.ProfileCfg, error)
	GetProfileCfgByScheduler(context.Context, dto.ProfileCfgScheduler) (*[]model.ProfileCfg, error)
}
