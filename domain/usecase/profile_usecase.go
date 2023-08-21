package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . ProfileUsecase
type ProfileUsecase interface {
	GetProfileById(context.Context, string) (*dto.ProfileResp, error)
	// StoreProfile(context.Context, string) (*Profile, error)
}
