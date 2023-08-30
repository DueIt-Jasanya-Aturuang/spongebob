package unit

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/validation"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileCfgDTO(t *testing.T) {
	t.Run("SUCCESS_ProfileCfgDTO_CREATE", func(t *testing.T) {
		reqCreate := dto.CreateProfileCfgReq{
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ProfileID:   "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ConfigValue: "19:00 Asia/Jakarta",
			Days:        []string{"monday", "tuesday"},
			ConfigName:  "DAILY_NOTIFY",
			Status:      "on",
			Token:       "12345678901",
		}

		err := validation.CreateProfileCfg(&reqCreate)
		assert.NoError(t, err)
	})

	t.Run("ERROR_ProfileCfgDTO_CREATE", func(t *testing.T) {
		reqCreate := dto.CreateProfileCfgReq{
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ProfileID:   "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ConfigValue: "as",
			Days:        []string{"mondays", "tuesday"},
			ConfigName:  "DAILY_NOTIFY",
			Status:      "osn",
			Token:       "678901",
		}

		err := validation.CreateProfileCfg(&reqCreate)
		t.Log(err)
		assert.Error(t, err)
	})

	t.Run("SUCCESS_ProfileCfgDTO_UPDATE", func(t *testing.T) {
		reqUpdate := dto.UpdateProfileCfgReq{
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ProfileID:   "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ConfigValue: "19:00 Asia/Jakarta",
			Days:        []string{"monday", "tuesday"},
			Status:      "on",
			Token:       "12345678901",
			ConfigName:  "DAILY_NOTIFY",
		}
		err := validation.UpdateProfileCfgValidate(&reqUpdate)
		assert.NoError(t, err)
	})

	t.Run("ERROR_ProfileCfgDTO_UPDATE", func(t *testing.T) {
		reqUpdate := dto.UpdateProfileCfgReq{
			UserID:      "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ProfileID:   "699137ef-1f24-46d7-82bf-862fde7b36d8",
			ConfigValue: "1900",
			Days:        []string{"mondays", "tuesday"},
			Status:      "osn",
			Token:       "678901",
			ConfigName:  "asd",
		}

		err := validation.UpdateProfileCfgValidate(&reqUpdate)
		assert.Error(t, err)
	})
}
