package format

import (
	"fmt"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"github.com/rs/zerolog/log"
)

func FormatConfigValue(configName, value, ianaTimezone string, days []string) (map[string]any, error) {
	configValue := map[string]any{}

	if configName == "DAILY_NOTIFY" {
		layout, err := time.Parse("15:04", value)
		if err != nil {
			log.Info().Msgf("error : %s | value : %s", exception.Err400InvalidTimeLayout.Error(), value)
			return nil, exception.Err400InvalidTimeLayout
		}

		loc, err := time.LoadLocation(ianaTimezone)
		if err != nil {
			log.Info().Msgf("error : %s | value : %s", exception.Err400InvalidIanaTimezone.Error(), ianaTimezone)
			return nil, exception.Err400InvalidIanaTimezone
		}

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
