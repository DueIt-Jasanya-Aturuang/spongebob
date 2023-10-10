package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func EmailFormat(email string) string {
	emailArr := strings.Split(email, "@")
	if len(emailArr) != 2 {
		log.Err(fmt.Errorf("%s", "invalid email user")).Msgf("INVALID EMAIL : %s", email)
		return email
	}
	return fmt.Sprintf("%c••••%c@%s", emailArr[0][0], emailArr[0][len(emailArr[0])-1], emailArr[1])
}

func ConfigValue(configName, value, ianaTimezone string, days []string) (map[string]any, error) {
	configValue := map[string]any{}

	if configName == "DAILY_NOTIFY" {
		layout, _ := time.Parse("15:04", value)

		loc, _ := time.LoadLocation(ianaTimezone)

		timeLayout := time.Date(2006, 0o1, 0o2, layout.Hour(), layout.Minute(), 0, 0, loc)

		configValue["config_time_user"] = value
		configValue["config_timezone_user"] = ianaTimezone
		configValue["config_time_notify"] = fmt.Sprintf("%02d:%02d", timeLayout.UTC().Hour(), timeLayout.UTC().Minute())
		configValue["config_timezone_notify"] = "UTC"
		configValue["days"] = days
	} else if configName == "MONTHLY_PERIOD" {
		configValue["config_date"] = value
	}

	return configValue, nil
}

func GetOrder(order string) (string, string) {
	var orderRes string
	if order != "asc" && order != "desc" {
		orderRes = "desc"
	} else {
		orderRes = order
	}

	var operation string
	if orderRes == "asc" {
		operation = ">"
	} else {
		operation = "<"
	}

	return orderRes, operation
}
