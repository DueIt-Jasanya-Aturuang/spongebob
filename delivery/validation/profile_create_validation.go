package validation

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

func StoreProfileValidation(req *dto.StoreProfileReq) error {
	if req.UserID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}

	return nil
}
