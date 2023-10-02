package converter

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase/helpers"
)

func UpdateAccountToModel(req *domain.RequestUpdateAccount, image string) (*domain.Profile, *domain.User) {
	timeUnix := time.Now().Unix()
	profile := &domain.Profile{
		ProfileID: req.ProfileID,
		UserID:    req.UserID,
		Quote:     helpers.NewNullString(req.Quote),
		Profesi:   helpers.NewNullString(req.Profesi),
		AuditInfo: domain.AuditInfo{
			UpdatedAt: timeUnix,
			UpdatedBy: helpers.NewNullString(req.ProfileID),
		},
	}

	user := &domain.User{
		ID:          req.UserID,
		FullName:    req.FullName,
		Gender:      req.Gender,
		Image:       image,
		PhoneNumber: helpers.NewNullString(req.PhoneNumber),
		AuditInfo: domain.AuditInfo{
			UpdatedAt: timeUnix,
			UpdatedBy: helpers.NewNullString(req.UserID),
		},
	}

	return profile, user
}

func UpdateAccountModelToResp(u *domain.User, p *domain.Profile, emailFormat string) (*domain.ResponseUser, *domain.ResponseProfile) {
	user := &domain.ResponseUser{
		ID:              u.ID,
		FullName:        u.FullName,
		Gender:          u.Gender,
		Image:           u.Image,
		Username:        u.Username,
		Email:           u.Email,
		EmailFormat:     emailFormat,
		PhoneNumber:     helpers.GetNullString(u.PhoneNumber),
		EmailVerifiedAt: false,
	}

	profile := &domain.ResponseProfile{
		ProfileID: p.ProfileID,
		Quote:     helpers.GetNullString(p.Quote),
		Profesi:   helpers.GetNullString(p.Profesi),
	}

	return user, profile
}

func ProfileDefault(userID string) *domain.Profile {
	id := uuid.NewV4().String()
	return &domain.Profile{
		ProfileID: id,
		UserID:    userID,
		Quote:     sql.NullString{},
		Profesi:   sql.NullString{},
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: id,
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}
}

func ProfileModelToResp(p *domain.Profile) *domain.ResponseProfile {
	return &domain.ResponseProfile{
		ProfileID: p.ProfileID,
		Quote:     helpers.GetNullString(p.Quote),
		Profesi:   helpers.GetNullString(p.Profesi),
	}
}
