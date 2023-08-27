package validation

import (
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils"
	"strconv"
	"strings"
)

func UpdateProfileCfgValidate(req *dto.UpdateProfileCfgReq) error {
	badReq := map[string][]string{}

	if req.UserID == "" {
		return exception.Err401(exception.Err401Msg.Error())
	}
	if len(req.UserID) > 40 {
		return exception.Err401(exception.Err401Msg.Error())
	}
	if len(req.UserID) < 30 {
		return exception.Err401(exception.Err401Msg.Error())
	}

	// profile_id validation
	if req.ProfileID == "" {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.Required, "profile_id"))
	}
	if len(req.ProfileID) > 50 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MaxString, "profile_id", 50))
	}
	if len(req.ProfileID) < 5 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MinString, "profile_id", 5))
	}

	// config_value validation
	// config_name validation
	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		badReq["config_name"] = append(badReq["config_name"], fmt.Sprintf(exception.InvalidField, "config_name", "your input", "DAILY_NOTIFY or MONTHLY_PERIOD"))
	}

	// config_value validation
	if req.ConfigValue == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.Required, "config_value"))
	}
	if req.ConfigName == "DAILY_NOTIFY" && len(strings.Split(req.ConfigValue, " ")) != 2 {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
	}
	if req.ConfigName == "DAILY_NOTIFY" && len(strings.Split(req.ConfigValue, " ")) == 2 {
		if strings.Split(req.ConfigValue, " ")[0] == "" || strings.Split(req.ConfigValue, " ")[1] == "" {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
		}
	}
	if req.ConfigName == "MONTHLY_PERIOD" {
		configValueInt, err := strconv.Atoi(req.ConfigValue)
		if err != nil {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.Integer, "config_value"))
		} else {
			if configValueInt > 29 {
				badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.MaxInteger, "config_value", 29))
			}
			if configValueInt < 1 {
				badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.MinInteger, "config_value", 1))
			}
		}
	}

	// days validation
	if !dayValidate(req.Days) {
		badReq["days"] = append(badReq["days"], fmt.Sprintf(exception.InvalidField, "days", "your input", strings.Join(utils.Days(), ", ")))
	}

	// status validation
	if req.Status != "on" && req.Status != "off" {
		badReq["status"] = append(badReq["status"], fmt.Sprintf(exception.Enum, "status", strings.Join([]string{"on", "off"}, " or ")))
	}

	// token validation
	if req.Token == "" {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.Required, "token"))
	}
	if len(req.Token) < 10 {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.MinInteger, "token", 10))
	}

	if len(badReq) >= 1 {
		return exception.Err400(badReq)
	}
	return nil
}
