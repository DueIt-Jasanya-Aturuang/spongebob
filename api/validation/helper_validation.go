package validation

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func dayValidate(days []string) bool {
	var status bool

	for _, configVal := range days {
		for _, v := range util.Days() {
			if v == configVal {
				status = true
				break
			}
			status = false
		}
		if !status {
			return false
		}
	}
	return true
}

const image = "image"

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

var (
	Required     = "%s is required"
	MaxString    = "maximum %s character must be %d"
	MinString    = "minimum %s character must be %d"
	Gender       = "%s gender must be male, female, or undefinied"
	FileSize     = "max %s size should be %d kb or %d mb"
	FileContent  = "%s must be %s"
	InvalidField = "%s invalid %s, example %s"
	Integer      = "%s must be number"
	Enum         = "%s must be %s"
	MinInteger   = "minimum %s number must be %d"
	MaxInteger   = "maximum %s number must be %d"
)
