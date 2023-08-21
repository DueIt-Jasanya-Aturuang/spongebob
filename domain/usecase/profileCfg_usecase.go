package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileCfgUsecase
type ProfileCfgUsecase interface {
	CreateProfileCfg(context.Context, dto.ProfileCfgReq) (*dto.ProfileCfgResp, error)
	GetProfileCfgById(context.Context, string) (*dto.ProfileCfgResp, error)
	UpdateProfileCfg(context.Context, dto.ProfileCfgReq, string) (*dto.ProfileCfgResp, error)
}
