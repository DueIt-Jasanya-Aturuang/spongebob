package validation

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/api/rest/response"
)

func GetProfileCfgValidation(req *dto.GetProfileCfgReq) error {
	if req.UserID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if req.ProfileID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.ProfileID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.ProfileID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}

	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		return response.Err404("NOT FOUND", nil)
	}

	return nil
}
