package dto

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
