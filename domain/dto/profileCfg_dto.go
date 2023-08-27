package dto

import "C"

// CreateProfileCfgReq create profile config request
type CreateProfileCfgReq struct {
	ProfileID    string   `json:"profile_id"`   // request body
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	ConfigName   string   `json:"config_name"`  // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	UserID       string   // request header
	Value        string   // helper
	IanaTimezone string   // helper
}

// UpdateProfileCfgReq update profile config request
type UpdateProfileCfgReq struct {
	ProfileID    string   `json:"profile_id"`   // request body
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	UserID       string   // request header
	ConfigName   string   // url parameter
	Value        string   // helper
	IanaTimezone string   // helper
}

// GetProfileCfgReq get profile config request
type GetProfileCfgReq struct {
	UserID     string // request header
	ConfigName string // url parameter config_name
	ProfileID  string // helper
}

type ProfileCfgResp struct {
	ID          string `json:"profile_config_id"`
	ProfileID   string `json:"profile_id"`
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
	Status      string `json:"status"`
}

type ProfileCfgScheduler struct {
	Day  string
	Time string
}
