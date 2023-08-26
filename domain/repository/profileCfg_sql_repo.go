package repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileCfgRepo
type ProfileCfgRepo interface {
	StoreProfileCfg(c context.Context, profileCfg model.ProfileCfg) error
	UpdateProfileCfg(c context.Context, profileCfg model.ProfileCfg) error
	GetProfileCfgByNameAndID(c context.Context, id, profileID, configName string) (*model.ProfileCfg, error)
	GetProfileCfgByScheduler(c context.Context, profileCfgScheduler dto.ProfileCfgScheduler) (*[]model.ProfileCfg, error)
	UoW() UnitOfWork
}
