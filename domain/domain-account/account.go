package domainaccount

import (
	"context"

	domainprofile "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile"
	domainuser "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-user"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./../mocks . AccountUsecase
type AccountUsecase interface {
	AccountUpdate(context.Context, UpdateAccountReq) (*domainuser.UserResp, *domainprofile.ProfileResp, error)
}
