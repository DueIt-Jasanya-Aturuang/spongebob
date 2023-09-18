package validation

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
)

func GetProfileCfgValidation(req *domain.RequestGetProfileConfig) error {
	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
	}

	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		return _error.HttpErrString("invalid config name", response.CM01)
	}

	return nil
}

func CreateProfileCfg(req *domain.RequestCreateProfileConfig) error {
	errBadRequest := map[string][]string{}

	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
	}

	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		errBadRequest["config_name"] = append(errBadRequest["config_name"], invalidConfig)
	}

	if req.ConfigValue == "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], required)
	}

	msg := checkConfigValue(req.ConfigName, req.ConfigValue)
	if msg != "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], msg)
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		msg = dayValidate(req.Days)
		if msg != "" {
			errBadRequest["days"] = append(errBadRequest["days"], msg)
		}
	}

	// status validation
	if req.Status != "on" && req.Status != "off" {
		errBadRequest["status"] = append(errBadRequest["status"], fmt.Sprintf(enum, "on", "off"))
	}

	// token validation
	if req.Token == "" {
		errBadRequest["token"] = append(errBadRequest["token"], required)
	}
	if len(req.Token) < 10 {
		errBadRequest["token"] = append(errBadRequest["token"], fmt.Sprintf(minNumeric, 10))
	}

	if len(errBadRequest) >= 1 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}

	return nil
}

func UpdateProfileCfgValidate(req *domain.RequsetUpdateProfileConfig) error {
	errBadRequest := map[string][]string{}

	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
	}

	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		errBadRequest["config_name"] = append(errBadRequest["config_name"], invalidConfig)
	}

	if req.ConfigValue == "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], required)
	}

	msg := checkConfigValue(req.ConfigName, req.ConfigValue)
	if msg != "" {
		errBadRequest["config_value"] = append(errBadRequest["config_value"], msg)
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		msg = dayValidate(req.Days)
		if msg != "" {
			errBadRequest["days"] = append(errBadRequest["days"], msg)
		}
	}

	if req.Status != "on" && req.Status != "off" {
		errBadRequest["status"] = append(errBadRequest["status"], fmt.Sprintf(enum, "on", "off"))
	}

	// token validation
	if req.Token == "" {
		errBadRequest["token"] = append(errBadRequest["token"], required)
	}
	if len(req.Token) < 10 {
		errBadRequest["token"] = append(errBadRequest["token"], fmt.Sprintf(minNumeric, 10))
	}

	if len(errBadRequest) >= 1 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}

	return nil
}
