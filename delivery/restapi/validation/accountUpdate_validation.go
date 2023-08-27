package validation

import (
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/exception"
	"strings"
)

// Validate account update request

func UpdateAccountValidate(req *dto.UpdateAccountReq) error {
	badReq := map[string][]string{}

	if req.UserID == "" {
		return exception.Err401(exception.Err401Msg.Error())
	}
	if len(req.UserID) > 40 {
		return exception.Err401(exception.Err401Msg.Error())
	}
	if len(req.UserID) < 30 {
		return exception.Err401(exception.Err401Msg.Error())
	}

	if req.ProfileID == "" {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.Required, "profile_id"))
	}
	if len(req.ProfileID) > 40 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.InvalidID, "profile_id"))
	}
	if len(req.ProfileID) < 30 {
		badReq["profile_id"] = append(badReq["profile_id"], fmt.Sprintf(exception.InvalidID, "profile_id"))
	}

	// fullName validation
	if req.FullName == "" {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.Required, "full_name"))
	}
	if len(req.FullName) > 32 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.MaxString, "full_name", 32))
	}
	if len(req.FullName) < 3 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(exception.MinString, "full_name", 3))
	}

	if req.Gender != "male" && req.Gender != "female" && req.Gender != "undefinied" {
		badReq["gender"] = append(badReq["gender"], fmt.Sprintf(exception.Gender, "gender"))
	}

	// image validation
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(exception.FileSize, "image", 2048, 2))
		}
		if !checkContentType(req.Image.Header.Get("Content-Type"), image) {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(exception.FileContent, "image", strings.Join(contentType(image), " or ")))
		}
	}

	// phoneNumber validation
	if req.PhoneNumber == "" {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.Required, "phone_number"))
	}
	if len(req.PhoneNumber) > 12 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.MaxString, "phone_number", 12))
	}
	if len(req.PhoneNumber) < 8 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(exception.MinString, "phone_number", 8))
	}

	// quote validation
	if req.Quote == "" {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.Required, "quote"))
	}
	if len(req.Quote) > 40 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.MaxString, "quote", 42))
	}
	if len(req.Quote) < 6 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(exception.MinString, "quote", 6))
	}

	if len(badReq) >= 1 {
		return exception.Err400(badReq)
	}

	return nil
}
