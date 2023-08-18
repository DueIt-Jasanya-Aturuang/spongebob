package domain

import (
	"context"
	"database/sql"
)

type Profile struct {
	ProfileId string         `json:"profile_id"`
	UserId    string         `json:"user_id"`
	Quote     sql.NullString `json:"quote"`
	CreatedAt int64          `json:"created_at"`
	CreatedBy string         `json:"created_by"`
	UpdatedAt int64          `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedAt sql.NullInt64  `json:"deleted_at"`
	DeletedBy sql.NullString `json:"deleted_by"`
}

type ProfileResp struct {
	ProfileID string  `json:"profile_id"`
	Quote     *string `json:"quote"`
}

type CreateProfileReq struct {
	UserId string
	Quote  string `json:"quote" form:"quote" validate:"required,min=6,max=128"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./mocks . ProfileRepo
type ProfileRepo interface {
	GetProfileById(context.Context, *sql.DB, string) (*Profile, error)
	GetProfileByUserId(context.Context, *sql.DB, string) (*Profile, error)
	StoreProfile(context.Context, *sql.Tx, Profile) (*Profile, error)
	UpdateProfile(context.Context, *sql.Tx, Profile) (*Profile, error)
}
