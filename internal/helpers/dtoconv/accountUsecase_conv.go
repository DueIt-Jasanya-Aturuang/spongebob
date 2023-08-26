package dtoconv

import (
	"database/sql"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

func UpdateAccountToModel(
	req dto.UpdateAccountReq,
	profileID, image string,
) (model.Profile, model.User) {
	timeUnix := time.Now().Unix()
	profile := model.Profile{
		ProfileID: profileID,
		UserID:    req.UserID,
		Quote:     sql.NullString{String: req.Quote},
		UpdatedAt: timeUnix,
		UpdatedBy: sql.NullString{String: profileID},
	}

	user := model.User{
		ID:          req.UserID,
		FullName:    req.FullName,
		Gender:      req.Gender,
		Image:       image,
		PhoneNumber: sql.NullString{String: req.PhoneNumber},
		UpdatedAt:   timeUnix,
		UpdatedBy:   sql.NullString{String: req.UserID},
	}

	return profile, user
}
