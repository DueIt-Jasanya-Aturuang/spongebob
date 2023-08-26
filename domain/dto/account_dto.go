package dto

import (
	"mime/multipart"
)

type UpdateAccountReq struct {
	UserID      string
	FullName    string                `json:"full_name" form:"full_name" validate:"required,min=3,max=32"`
	Gender      string                `json:"gender" form:"gender"`
	Image       *multipart.FileHeader `json:"image" form:"image"`
	PhoneNumber string                `json:"phone_number"`
	Quote       string                `json:"quote" form:"quote" validate:"required,min=6,max=128"`
}
