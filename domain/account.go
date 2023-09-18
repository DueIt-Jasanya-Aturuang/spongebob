package domain

import (
	"context"
	"mime/multipart"
)

type RequestUpdateAccount struct {
	UserID      string                // request header 'User-Id'
	ProfileID   string                // request param 'profile-id'
	FullName    string                `json:"full_name" form:"full_name"`       // request body
	Gender      string                `json:"gender" form:"gender"`             // request body
	Image       *multipart.FileHeader `json:"image" form:"image"`               // request body
	PhoneNumber string                `json:"phone_number" form:"phone_number"` // request body
	Quote       string                `json:"quote" form:"quote"`               // request body
	Profesi     string                `json:"profesi" form:"profesi"`           // request body
}

type RequestGetProfile struct {
	UserID string // request header 'id'
}

type RequestCreateProfile struct {
	UserID string `json:"user_id"`
}

type ResponseProfile struct {
	ProfileID string `json:"profile_id"`
	Quote     string `json:"quote"`
	Profesi   string `json:"profesi"`
}

type ResponseUser struct {
	ID              string `json:"id"`
	FullName        string `json:"full_name"`
	Gender          string `json:"gender"`
	Image           string `json:"image"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailFormat     string `json:"email_format"`
	PhoneNumber     string `json:"phone_number"`
	EmailVerifiedAt bool   `json:"activated"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ./../mocks . AccountUsecase
type AccountUsecase interface {
	UpdateAccount(ctx context.Context, req *RequestUpdateAccount) (*ResponseUser, *ResponseProfile, error)
	GetProfileByID(ctx context.Context, req *RequestGetProfile) (*ResponseProfile, error)
	CreateProfile(ctx context.Context, req *RequestCreateProfile) (*ResponseProfile, error)
}
