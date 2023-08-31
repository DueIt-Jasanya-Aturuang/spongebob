package helpers

import (
	"database/sql"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
)

func UpdateAccountToModel(
	req *dto.UpdateAccountReq,
	image string,
) (model.Profile, model.User) {
	timeUnix := time.Now().Unix()
	profile := model.Profile{
		ProfileID: req.ProfileID,
		UserID:    req.UserID,
		Quote:     sql.NullString{String: req.Quote},
		Profesi:   sql.NullString{String: req.Profesi},
		UpdatedAt: timeUnix,
		UpdatedBy: sql.NullString{String: req.ProfileID},
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
