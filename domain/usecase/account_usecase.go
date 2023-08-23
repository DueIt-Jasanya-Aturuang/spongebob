package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ./../mocks . AccountUsecase
type AccountUsecase interface {
	AccountUpdate(context.Context, dto.UpdateAccountReq) (*dto.UserResp, *dto.ProfileResp, error)
}
