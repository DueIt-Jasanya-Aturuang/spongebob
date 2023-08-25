package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//counterfeiter:generate -o ./../mocks . ProfileCfgUsecase
type ProfileCfgUsecase interface {
	CreateProfileCfg(context.Context, dto.CreateProfileCfgReq) (*dto.ProfileCfgResp, error)
	GetProfileCfgByNameAndID(c context.Context, id, profileID, configName string) (*dto.ProfileCfgResp, error)
	UpdateProfileCfg(c context.Context, req dto.UpdateProfileCfgReq, id string, configName string) (*dto.ProfileCfgResp, error)
}
