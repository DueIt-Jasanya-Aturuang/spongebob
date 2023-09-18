package helpers

import (
	"fmt"
	"time"
)

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
