package helpers

import (
	"errors"
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"time"

	"github.com/rs/zerolog/log"
)

func ConfigValue(configName, value, ianaTimezone string, days []string) (map[string]any, error) {
	configValue := map[string]any{}

	if configName == "DAILY_NOTIFY" {
		layout, err := time.Parse("15:04", value)
		if err != nil {
			log.Err(errors.New(message.ErrInvalidTimeLayout)).Msgf("error : %s | value : %s", message.ErrInvalidTimeLayout, value)
			return nil, model.ErrBadInput
		}

		loc, err := time.LoadLocation(ianaTimezone)
		if err != nil {
			log.Err(errors.New(message.ErrInvalidIanaTimezone)).Msgf("error : %s | value : %s", message.ErrInvalidIanaTimezone, ianaTimezone)
			return nil, model.ErrBadInput
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
