package domain

import "database/sql"

type ProfileConfig struct {
	ID          string         `json:"user_config_id"`
	ProfileId   string         `json:"profile_id"`
	ConfigName  string         `json:"config_name"`
	ConfigValue string         `json:"config_value"`
	Status      string         `json:"status"`
	CreatedAt   int64          `json:"created_at"`
	CreatedBy   string         `json:"created_by"`
	UpdatedAt   int64          `json:"updated_at"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	DeletedAt   sql.NullInt64  `json:"deleted_at"`
	DeletedBy   sql.NullString `json:"deleted_by"`
}

type ProfileConfigReq struct {
	ProfileID   string   `json:"profile_id" validate:"required"`
	ConfigValue string   `json:"config_value" validate:"required"`
	Days        []string `json:"days"`
	ConfigName  string   `json:"config_name" validate:"required"`
	Status      string   `json:"status" validate:"required"`
	Token       string   `json:"token" validate:"required"`
}

type ProfileConfigResp struct {
	ID          string `json:"user_config_id"`
	ProfileId   string `json:"profile_id"`
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
	Status      string `json:"status"`
}
