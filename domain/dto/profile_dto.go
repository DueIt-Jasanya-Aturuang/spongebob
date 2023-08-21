package dto

import "database/sql"

type ProfileResp struct {
	ProfileID string         `json:"profile_id"`
	Quote     sql.NullString `json:"quote"`
}

type CreateProfileReq struct {
	UserId string
	Quote  string `json:"quote" form:"quote" validate:"required,min=6,max=128"`
}
