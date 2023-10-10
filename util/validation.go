package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const Image = "image"

var Required = "field ini tidak boleh dikosongkan"
var MinString = "field ini tidak boleh kurang dari %d"
var MaxString = "field ini tidak boleh lebih dari %d"
var Enum = "harus %s atau %s"
var MinNumeric = "field ini tidak boleh kurang dari %d"
var MaxNumeric = "field ini tidak boleh lebih dari %d"
var FileSize = "maximal size harus %d kb atau %d mb"
var FileContent = "file content harus %s"
var InvalidConfig = "invalid config name"
var InvalidConfigValue = "invalid config value, contoh %s"

func DayValidate(days []string) string {
	var status bool

	if len(days) < 1 {
		return "setidaknya harus memilih 1 hari"
	}

	for _, configVal := range days {
		for _, v := range Days() {
			if v == configVal {
				status = true
				break
			}
			status = false
		}
		if !status {
			return "nama hari tidak sesuai format"
		}
	}

	return ""
}

func ContentType(typeContent string) []string {
	switch typeContent {
	case Image:
		return []string{
			"image/png", "image/jpeg", "image/jpg",
		}
	}

	return []string{
		"image/png", "image/jpeg", "image/jpg",
	}
}

func CheckContentType(headerContentType string, typeContent string) bool {
	if headerContentType == "" {
		return false
	}

	var status bool
	for _, v := range ContentType(typeContent) {
		if headerContentType == v {
			return true
		}
		status = false
	}
	return status
}

func CheckConfigValue(name, value string) string {
	if name == "DAILY_NOTIFY" {
		valueSplit := strings.Split(value, " ")
		if len(valueSplit) != 2 {
			return fmt.Sprintf(InvalidConfigValue, "19:20 Asia/Jakarta")
		}

		_, err := time.Parse("15:04", valueSplit[0])
		if err != nil {
			log.Warn().Msgf("failed parse time : %v", err)
			return fmt.Sprintf(InvalidConfigValue, "19:20 Asia/Jakarta")
		}

		_, err = time.LoadLocation(valueSplit[1])
		if err != nil {
			log.Warn().Msgf("failed load location : %v", err)
			return fmt.Sprintf(InvalidConfigValue, "19:20 Asia/Jakarta")
		}
	} else {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return "config value harus integer"
		}

		if valueInt > 29 {
			return fmt.Sprintf(MaxNumeric, 29)
		}
		if valueInt < 1 {
			return fmt.Sprintf(MinNumeric, 1)
		}
	}

	return ""
}

func MaxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(MinString, min)
	case len(s) > max:
		return fmt.Sprintf(MaxString, max)
	}

	return ""
}
