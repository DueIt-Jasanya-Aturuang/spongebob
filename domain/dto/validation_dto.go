package dto

import (
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils"
	"strconv"
	"strings"
)

// Validate account update request
func (req *UpdateAccountReq) contentType() []string {
	return []string{
		"image/png", "image/jpeg", "image/jpg",
	}
}
func (req *UpdateAccountReq) checkContentType(headerContentType string) bool {
	if headerContentType == "" {
		return false
	}

	var status bool
	for _, v := range req.contentType() {
		if headerContentType == v {
			return true
		}
		status = false
	}
	return status
}
func (req *UpdateAccountReq) Validate() error {
	badReq := map[string][]string{}

	if req.ProfileID == "" {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.Required, "profile_id"))
	}
	// fullName validation
	if req.FullName == "" {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.Required, "full_name"))
	}
	if len(req.FullName) > 32 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.MaxString, "full_name", 32))
	}
	if len(req.FullName) < 3 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.MinString, "full_name", 3))
	}

	if req.Gender != "male" && req.Gender != "female" && req.Gender != "undefinied" {
		badReq["gender"] = append(badReq["gender"], fmt.Sprintf(exception.Gender, "gender"))
	}

	// image validation
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(exception.FileSize, "image", 2048, 2))
		}
		if !req.checkContentType(req.Image.Header.Get("Content-Type")) {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(exception.FileContent, "image", strings.Join(req.contentType(), " or ")))
		}
	}

	// phoneNumber validation
	if req.PhoneNumber == "" {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.Required, "phone_number"))
	}
	if len(req.PhoneNumber) > 12 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.MaxString, "phone_number", 12))
	}
	if len(req.PhoneNumber) < 8 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.MinString, "phone_number", 8))
	}

	// quote validation
	if req.Quote == "" {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.Required, "quote"))
	}
	if len(req.Quote) > 40 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.MaxString, "quote", 42))
	}
	if len(req.Quote) < 6 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.MinString, "quote", 6))
	}

	if len(badReq) >= 1 {
		return exception.Err400(badReq)
	}

	return nil
}

// Validate create profile config request
func (c *CreateProfileCfgReq) Validate() error {
	badReq := map[string][]string{}

	// profile_id validation
	if c.ProfileID == "" {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.Required, "profile_id"))
	}
	if len(c.ProfileID) > 50 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MaxString, "profile_id", 50))
	}
	if len(c.ProfileID) < 5 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MinString, "profile_id", 5))
	}

	// config_name validation
	if c.ConfigName != "DAILY_NOTIFY" && c.ConfigName != "MONTHLY_PERIOD" {
		badReq["config_name"] = append(badReq["config_name"], fmt.Sprintf(exception.InvalidField, "config_name", "your input", "DAILY_NOTIFY or MONTHLY_PERIOD"))
	}

	// config_value validation
	if c.ConfigValue == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.Required, "config_value"))
	}
	if c.ConfigName == "DAILY_NOTIFY" && len(strings.Split(c.ConfigValue, " ")) != 2 || strings.Split(c.ConfigValue, " ")[0] == "" && strings.Split(c.ConfigValue, " ")[1] == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
	}
	if c.ConfigName == "MONTHLY_PERIOD" {
		configValueInt, err := strconv.Atoi(c.ConfigValue)
		if err != nil {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.Integer, "config_value"))
		}

		if configValueInt > 29 {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.MaxInteger, "config_value", 29))
		}
		if configValueInt < 1 {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.MinInteger, "config_value", 1))
		}
	}

	// days validation
	if !dayValidate(c.Days) {
		badReq["days"] = append(badReq["days"], fmt.Sprintf(exception.InvalidField, "days", "your input", strings.Join(utils.Days(), ", ")))
	}
	if len(badReq) >= 1 {
		return exception.Err400(badReq)
	}

	// status validation
	if c.Status != "on" && c.Status != "off" {
		badReq["status"] = append(badReq["status"], fmt.Sprintf(exception.Enum, "status", strings.Join([]string{"on", "off"}, " or ")))
	}

	// token validation
	if c.Token == "" {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.Required, "token"))
	}
	if len(c.Token) < 10 {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.MinInteger, "token", 10))
	}
	return nil
}

// Validate update profile config request
func (u *UpdateProfileCfgReq) Validate(configName string) error {
	badReq := map[string][]string{}

	// profile_id validation
	if u.ProfileID == "" {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.Required, "profile_id"))
	}
	if len(u.ProfileID) > 50 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MaxString, "profile_id", 50))
	}
	if len(u.ProfileID) < 5 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.MinString, "profile_id", 5))
	}

	// config_value validation
	// config_name validation
	if configName != "DAILY_NOTIFY" && configName != "MONTHLY_PERIOD" {
		badReq["config_name"] = append(badReq["config_name"], fmt.Sprintf(exception.InvalidField, "config_name", "your input", "DAILY_NOTIFY or MONTHLY_PERIOD"))
	}

	// config_value validation
	if u.ConfigValue == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.Required, "config_value"))
	}
	if configName == "DAILY_NOTIFY" && len(strings.Split(u.ConfigValue, " ")) != 2 || strings.Split(u.ConfigValue, " ")[0] == "" && strings.Split(u.ConfigValue, " ")[1] == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(exception.InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
	}
	if configName == "MONTHLY_PERIOD" {
		configValueInt, err := strconv.Atoi(u.ConfigValue)
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
	if !dayValidate(u.Days) {
		badReq["days"] = append(badReq["days"], fmt.Sprintf(exception.InvalidField, "days", "your input", strings.Join(utils.Days(), ", ")))
	}
	if len(badReq) >= 1 {
		return exception.Err400(badReq)
	}

	// status validation
	if u.Status != "on" && u.Status != "off" {
		badReq["status"] = append(badReq["status"], fmt.Sprintf(exception.Enum, "status", strings.Join([]string{"on", "off"}, " or ")))
	}

	// token validation
	if u.Token == "" {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.Required, "token"))
	}
	if len(u.Token) < 10 {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(exception.MinInteger, "token", 10))
	}
	return nil
}
