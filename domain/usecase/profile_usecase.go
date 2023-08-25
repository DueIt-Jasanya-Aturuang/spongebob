package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//counterfeiter:generate -o ./../mocks . ProfileUsecase
type ProfileUsecase interface {
	GetProfileByID(c context.Context, id string) (*dto.ProfileResp, error)
	// StoreProfile(context.Context, string) (*Profile, error)
}
