package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileCfgRepo
type ProfileCfgRepo interface {
	StoreProfileCfg(context.Context, model.ProfileCfg) error
	UpdateProfileCfg(context.Context, model.ProfileCfg) error
	GetProfileCfgByID(context.Context, string) (*model.ProfileCfg, error)
	GetProfileCfgByScheduler(context.Context, dto.ProfileCfgScheduler) (*[]model.ProfileCfg, error)
	UnitOfWork
}
