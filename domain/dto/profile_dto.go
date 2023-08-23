package dto

import "database/sql"

type ProfileResp struct {
	ProfileID string         `json:"profile_id"`
	Quote     sql.NullString `json:"quote"`
}
