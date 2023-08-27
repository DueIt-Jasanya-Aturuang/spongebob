package usecase

import (
	"context"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//counterfeiter:generate -o ./../mocks . ProfileUsecase
type ProfileUsecase interface {
	GetProfileByID(c context.Context, req dto.GetProfileReq) (*dto.ProfileResp, error)
	StoreProfile(c context.Context, req dto.StoreProfileReq) (*model.Profile, error)
}
