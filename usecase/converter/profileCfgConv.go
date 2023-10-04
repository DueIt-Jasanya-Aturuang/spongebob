package converter

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
)

func CreateProfileCfgToModel(req *domain.RequestCreateProfileConfig, configValue []byte) *domain.ProfileConfig {
	id := uuid.NewV4().String()
	return &domain.ProfileConfig{
		ID:          id,
		ProfileID:   req.ProfileID,
		ConfigName:  req.ConfigName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func UpdateProfileCfgToModel(req *domain.RequsetUpdateProfileConfig, configValue []byte, configName, id string) *domain.ProfileConfig {
	return &domain.ProfileConfig{
		ID:          id,
		ProfileID:   req.ProfileID,
		ConfigName:  configName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{String: req.ProfileID},
		},
	}
}

func ProfileConfigModelToResponse(m *domain.ProfileConfig, configValue string, days []string) *domain.ResponseProfileConfig {
	return &domain.ResponseProfileConfig{
		ID:          m.ID,
		ProfileID:   m.ProfileID,
		ConfigName:  m.ConfigName,
		ConfigValue: configValue,
		Status:      m.Status,
		Days:        days,
	}
}
