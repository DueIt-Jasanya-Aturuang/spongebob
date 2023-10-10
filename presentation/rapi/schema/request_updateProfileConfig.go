package schema

import (
	"fmt"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type RequsetUpdateProfileConfig struct {
	ConfigValue string   `json:"config_value"`
	Days        []string `json:"days"`
	Status      string   `json:"status"`
	Token       string   `json:"token"`
}

func (req *RequsetUpdateProfileConfig) Validate(configName string) error {
	errBadRequest := map[string][]string{}

	if configName != "DAILY_NOTIFY" && configName != "MONTHLY_PERIOD" {
		errBadRequest["config_name"] = append(errBadRequest["config_name"], util.InvalidConfig)
	}

	if req.ConfigValue == "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], util.Required)
	}

	msg := util.CheckConfigValue(configName, req.ConfigValue)
	if msg != "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], msg)
	}

	if configName == "DAILY_NOTIFY" {
		msg = util.DayValidate(req.Days)
		if msg != "" {
			errBadRequest["days"] = append(errBadRequest["days"], msg)
		}
	}

	if req.Status != "on" && req.Status != "off" {
		errBadRequest["status"] = append(errBadRequest["status"], fmt.Sprintf(util.Enum, "on", "off"))
	}

	// token validation
	if req.Token == "" {
		errBadRequest["token"] = append(errBadRequest["token"], util.Required)
	}
	if len(req.Token) < 10 {
		errBadRequest["token"] = append(errBadRequest["token"], fmt.Sprintf(util.MinNumeric, 10))
	}

	if len(errBadRequest) >= 1 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}

	return nil
}
