package domainaccount

import "mime/multipart"

type UpdateAccountReq struct {
	UserID      string
	FullName    string                `json:"full_name" form:"full_name" validate:"required,min=3,max=32"`
	Gender      string                `json:"gender" form:"gender"`
	Image       *multipart.FileHeader `json:"image" form:"image" swaggerignore:"true"`
	Username    string                `json:"username" form:"username" validate:"required,min=3,max=22"`
	PhoneNumber string
	Quote       string `json:"quote" form:"quote" validate:"required,min=6,max=128"`
}
