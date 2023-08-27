package unit

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileCfgDTO(t *testing.T) {
	t.Run("SUCCESS_ProfileCfgDTO_CREATE", func(t *testing.T) {
		reqCreate := dto.CreateProfileCfgReq{
			ProfileID:   "profileid1",
			ConfigValue: "19:00 Asia/Jakarta",
			Days:        []string{"monday", "tuesday"},
			ConfigName:  "DAILY_NOTIFY",
			Status:      "on",
			Token:       "12345678901",
		}

		err := reqCreate.Validate()
		assert.NoError(t, err)
	})

	t.Run("ERROR_ProfileCfgDTO_CREATE", func(t *testing.T) {
		reqCreate := dto.CreateProfileCfgReq{
			ProfileID:   "123",
			ConfigValue: "19:00 Asia/Jakarta",
			Days:        []string{"mondays", "tuesday"},
			ConfigName:  "DAILY_NOTIFY",
			Status:      "osn",
			Token:       "678901",
		}

		err := reqCreate.Validate()
		assert.Error(t, err)
	})

	t.Run("SUCCESS_ProfileCfgDTO_UPDATE", func(t *testing.T) {
		reqUpdate := dto.UpdateProfileCfgReq{
			ProfileID:   "profileid1",
			ConfigValue: "19:00 Asia/Jakarta",
			Days:        []string{"monday", "tuesday"},
			Status:      "on",
			Token:       "12345678901",
		}
		err := reqUpdate.Validate("DAILY_NOTIFY")
		assert.NoError(t, err)
	})

	t.Run("ERROR_ProfileCfgDTO_UPDATE", func(t *testing.T) {
		reqUpdate := dto.UpdateProfileCfgReq{
			ProfileID:   "123",
			ConfigValue: "1900",
			Days:        []string{"mondays", "tuesday"},
			Status:      "osn",
			Token:       "678901",
		}

		err := reqUpdate.Validate("MONTHLY_sPERIOD")
		assert.Error(t, err)
	})
}
