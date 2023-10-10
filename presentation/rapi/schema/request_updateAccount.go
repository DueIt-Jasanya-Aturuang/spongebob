package schema

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

type RequestUpdateAccount struct {
	FullName    string                `json:"full_name" form:"full_name"`       // request body
	Gender      string                `json:"gender" form:"gender"`             // request body
	Image       *multipart.FileHeader `json:"image" form:"image"`               // request body
	PhoneNumber string                `json:"phone_number" form:"phone_number"` // request body
	Quote       string                `json:"quote" form:"quote"`               // request body
	Profesi     string                `json:"profesi" form:"profesi"`           // request body
}

func (req *RequestUpdateAccount) Validate() error {
	errBadRequest := map[string][]string{}

	// fullName validation
	if req.FullName == "" {
		errBadRequest["full_name"] = append(errBadRequest["full_name"], util.Required)
	}
	fullName := util.MaxMinString(req.FullName, 3, 32)
	if fullName != "" {
		errBadRequest["full_name"] = append(errBadRequest["full_name"], fullName)
	}

	if req.Profesi != "" {
		profesi := util.MaxMinString(req.Profesi, 6, 30)
		if profesi != "" {
			errBadRequest["profesi"] = append(errBadRequest["profesi"], profesi)
		}
	}

	if req.Gender != "male" && req.Gender != "female" && req.Gender != "undefined" {
		errBadRequest["gender"] = append(errBadRequest["gender"], "invalid gender")
	}

	// image validation
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			errBadRequest["image"] = append(errBadRequest["image"], fmt.Sprintf(util.FileSize, 2048, 2))
		}
		if !util.CheckContentType(req.Image.Header.Get("Content-Type"), util.Image) {
			errBadRequest["image"] = append(errBadRequest["image"], fmt.Sprintf(util.FileContent, strings.Join(util.ContentType(util.Image), " or ")))
		}
	}

	// phoneNumber validation
	if req.PhoneNumber == "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], util.Required)
	}
	phoneNumber := util.MaxMinString(req.PhoneNumber, 8, 12)
	if phoneNumber != "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], phoneNumber)
	}

	// quote validation
	if req.Quote != "" {
		quote := util.MaxMinString(req.Quote, 6, 42)
		if quote != "" {
			errBadRequest["quote"] = append(errBadRequest["quote"], quote)
		}
	}

	if len(errBadRequest) >= 1 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}

	return nil
}
