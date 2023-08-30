package dto

import (
	"mime/multipart"
)

type UpdateAccountReq struct {
	UserID      string                // request header 'User-Id'
	ProfileID   string                // request param 'profile-id'
	FullName    string                `json:"full_name" form:"full_name"`       // request body
	Gender      string                `json:"gender" form:"gender"`             // request body
	Image       *multipart.FileHeader `json:"image" form:"image"`               // request body
	PhoneNumber string                `json:"phone_number" form:"phone_number"` // request body
	Quote       string                `json:"quote" form:"quote"`               // request body
}
