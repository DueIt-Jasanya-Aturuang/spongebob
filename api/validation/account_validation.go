package validation

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
)

func CreateProfileValidation(req *domain.RequestCreateProfile) error {
	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	return nil
}

func GetProfileValidation(req *domain.RequestGetProfile) error {
	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	return nil
}

func UpdateAccountValidate(req *domain.RequestUpdateAccount) error {
	errBadRequest := map[string][]string{}

	if _, err := uuid.Parse(req.UserID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04)
	}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
	}

	// fullName validation
	if req.FullName == "" {
		errBadRequest["full_name"] = append(errBadRequest["full_name"], required)
	}
	fullName := maxMinString(req.FullName, 3, 32)
	if fullName != "" {
		errBadRequest["full_name"] = append(errBadRequest["full_name"], fullName)
	}

	if req.Profesi != "" {
		profesi := maxMinString(req.Profesi, 6, 30)
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
			errBadRequest["image"] = append(errBadRequest["image"], fmt.Sprintf(fileSize, 2048, 2))
		}
		if !checkContentType(req.Image.Header.Get("Content-Type"), image) {
			errBadRequest["image"] = append(errBadRequest["image"], fmt.Sprintf(fileContent, strings.Join(contentType(image), " or ")))
		}
	}

	// phoneNumber validation
	if req.PhoneNumber == "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], required)
	}
	phoneNumber := maxMinString(req.PhoneNumber, 8, 12)
	if phoneNumber != "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], phoneNumber)
	}

	// quote validation
	if req.Quote != "" {
		quote := maxMinString(req.Quote, 6, 42)
		if quote != "" {
			errBadRequest["quote"] = append(errBadRequest["quote"], quote)
		}
	}

	if len(errBadRequest) >= 1 {
		return _error.HttpErrMapOfSlices(errBadRequest, response.CM06)
	}

	return nil
}
