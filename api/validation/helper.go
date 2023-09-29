package validation

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/rs/zerolog/log"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

const image = "image"

var required = "field ini tidak boleh dikosongkan"
var minString = "field ini tidak boleh kurang dari %d"
var maxString = "field ini tidak boleh lebih dari %d"
var enum = "harus %s atau %s"
var minNumeric = "field ini tidak boleh kurang dari %d"
var maxNumeric = "field ini tidak boleh lebih dari %d"
var fileSize = "maximal size harus %d kb atau %d mb"
var fileContent = "file content harus %s"
var invalidConfig = "invalid config name"
var invalidConfigValue = "invalid config value, contoh %s"

func dayValidate(days []string) string {
	var status bool

	if len(days) < 1 {
		return "setidaknya harus memilih 1 hari"
	}

	for _, configVal := range days {
		for _, v := range util.Days() {
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

// file validate
func contentType(typeContent string) []string {
	switch typeContent {
	case image:
		return []string{
			"image/png", "image/jpeg", "image/jpg",
		}
	}

	return []string{
		"image/png", "image/jpeg", "image/jpg",
	}
}

func checkContentType(headerContentType string, typeContent string) bool {
	if headerContentType == "" {
		return false
	}

	var status bool
	for _, v := range contentType(typeContent) {
		if headerContentType == v {
			return true
		}
		status = false
	}
	return status
}

func checkConfigValue(name, value string) string {
	if name == "DAILY_NOTIFY" {
		valueSplit := strings.Split(value, " ")
		if len(valueSplit) != 2 {
			return fmt.Sprintf(invalidConfigValue, "19:20 Asia/Jakarta")
		}

		_, err := time.Parse("15:04", valueSplit[0])
		if err != nil {
			log.Warn.Msgf("failed parse time : %v",err)
			return fmt.Sprintf(invalidConfigValue, "19:20 Asia/Jakarta")
		}

		_, err = time.LoadLocation(valueSplit[1])
		if err != nil {
			log.Warn.Msgf("failed load location : %v",err)
			return fmt.Sprintf(invalidConfigValue, "19:20 Asia/Jakarta")
		}
	} else {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return "config value harus integer"
		}

		if valueInt > 29 {
			return fmt.Sprintf(maxNumeric, 29)
		}
		if valueInt < 1 {
			return fmt.Sprintf(minNumeric, 1)
		}
	}

	return ""
}

func maxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(minString, min)
	case len(s) > max:
		return fmt.Sprintf(maxString, max)
	}

	return ""
}
