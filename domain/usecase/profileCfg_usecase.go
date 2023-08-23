package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)



//counterfeiter:generate -o ./../mocks . ProfileCfgUsecase
type ProfileCfgUsecase interface {
	CreateProfileCfg(context.Context, dto.ProfileCfgReq) (*dto.ProfileCfgResp, error)
	GetProfileCfgByID(context.Context, string) (*dto.ProfileCfgResp, error)
	UpdateProfileCfg(context.Context, dto.ProfileCfgReq, string) (*dto.ProfileCfgResp, error)
}
