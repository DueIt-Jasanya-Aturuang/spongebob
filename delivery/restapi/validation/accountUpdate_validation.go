package validation

import (
	"fmt"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"strings"
)

// Validate account update request

func UpdateAccountValidate(req *dto.UpdateAccountReq) error {
	badReq := map[string][]string{}

	if req.UserID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}

	if req.ProfileID == "" {
		return response.Err404("NOT FOUND", nil)
	}
	if len(req.ProfileID) > 36 {
		return response.Err404("NOT FOUND", nil)
	}
	if len(req.ProfileID) < 36 {
		return response.Err404("NOT FOUND", nil)
	}

	// fullName validation
	if req.FullName == "" {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(Required, "full_name"))
	}
	if len(req.FullName) > 32 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(MaxString, "full_name", 32))
	}
	if len(req.FullName) < 3 {
		badReq["full_name"] = append(badReq["full_name"], fmt.Sprintf(MinString, "full_name", 3))
	}

	if req.Gender != "male" && req.Gender != "female" && req.Gender != "undefinied" {
		badReq["gender"] = append(badReq["gender"], fmt.Sprintf(Gender, "gender"))
	}

	// image validation
	if req.Image != nil && req.Image.Size > 0 {
		if req.Image.Size > 2097152 {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(FileSize, "image", 2048, 2))
		}
		if !checkContentType(req.Image.Header.Get("Content-Type"), image) {
			badReq["image"] = append(badReq["image"], fmt.Sprintf(FileContent, "image", strings.Join(contentType(image), " or ")))
		}
	}

	// phoneNumber validation
	if req.PhoneNumber == "" {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(Required, "phone_number"))
	}
	if len(req.PhoneNumber) > 12 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(MaxString, "phone_number", 12))
	}
	if len(req.PhoneNumber) < 8 {
		badReq["phone_number"] = append(badReq["phone_number"], fmt.Sprintf(MinString, "phone_number", 8))
	}

	// quote validation
	if req.Quote == "" {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(Required, "quote"))
	}
	if len(req.Quote) > 40 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(MaxString, "quote", 42))
	}
	if len(req.Quote) < 6 {
		badReq["quote"] = append(badReq["quote"], fmt.Sprintf(MinString, "quote", 6))
	}

	if len(badReq) >= 1 {
		return response.Err400(badReq, nil)
	}

	return nil
}
