package util

import (
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

const DefaultMessageMonthlyPeriod = "default message"
const DefaultIconMonthlyPeriod = "default icon"
const DefaultTitleMonthlyPeriod = "default title"
const MonthlyPeriode = "MONTHLY_PERIOD"

const DefaultMessageDailyNotify = "default message"
const DefaultIconDailyNotify = "default icon"
const DefaultTitleDailyNotify = "default title"
const DailyNotify = "DAILY_NOTIFY"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var NewUlid = ulid.Make().String()

func ParseUlid(u string) error {
	if _, err := ulid.Parse(u); err != nil {
		log.Info().Msgf("failed parse ulid | err : %v", err)
		return err
	}

	return nil
}

func ParseUUID(u string) error {
	if _, err := ulid.Parse(u); err != nil {
		log.Info().Msgf("failed parse uuid | err : %v", err)
		return err
	}

	return nil
}

func Days() []string {
	return []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}
}
