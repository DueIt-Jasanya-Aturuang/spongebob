package dto

import "C"

// CreateProfileCfgReq create profile config request
type CreateProfileCfgReq struct {
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	ConfigName   string   `json:"config_name"`  // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	UserID       string   // request header
	ProfileID    string   // request param
	Value        string   // response
	IanaTimezone string   // response
}

// UpdateProfileCfgReq update profile config request
type UpdateProfileCfgReq struct {
	ConfigValue  string   `json:"config_value"` // request body
	Days         []string `json:"days"`         // request body
	Status       string   `json:"status"`       // request body
	Token        string   `json:"token"`        // request body
	ProfileID    string   // url parameter
	UserID       string   // request header
	ConfigName   string   // url parameter
	Value        string   // response
	IanaTimezone string   // response
}

// GetProfileCfgReq get profile config request
type GetProfileCfgReq struct {
	UserID     string // request header
	ConfigName string // url parameter config_name
	ProfileID  string // response
}

type ProfileCfgResp struct {
	ID          string `json:"profile_config_id"`
	ProfileID   string `json:"profile_id"`
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
	Status      string `json:"status"`
}

type ProfileCfgSche struct {
	Day  string
	Time string
}
