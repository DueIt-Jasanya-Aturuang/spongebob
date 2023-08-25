package dto

type ProfileCfgScheduler struct {
	Day  string
	Time string
}

type CreateProfileCfgReq struct {
	ProfileID    string   `json:"profile_id" validate:"required"`
	ConfigValue  string   `json:"config_value" validate:"required"`
	Days         []string `json:"days"`
	ConfigName   string   `json:"config_name" validate:"required"`
	Status       string   `json:"status" validate:"required"`
	Token        string   `json:"token" validate:"required"`
	Value        string
	IanaTimezone string
}

type UpdateProfileCfgReq struct {
	ProfileID    string   `json:"profile_id" validate:"required"`
	ConfigValue  string   `json:"config_value" validate:"required"`
	Days         []string `json:"days"`
	Status       string   `json:"status" validate:"required"`
	Token        string   `json:"token" validate:"required"`
	Value        string
	IanaTimezone string
}

type ProfileCfgResp struct {
	ID          string `json:"user_config_id"`
	ProfileID   string `json:"profile_id"`
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
	Status      string `json:"status"`
}
