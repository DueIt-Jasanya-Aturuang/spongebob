package validation

import "github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils"

func dayValidate(days []string) bool {
	var status bool

	for _, configVal := range days {
		for _, v := range utils.Days() {
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
