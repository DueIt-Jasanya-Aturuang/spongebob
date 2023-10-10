package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type AccountUsecase interface {
	UpdateAccount(ctx context.Context, req *RequestUpdateAccount) (*ResponseUser, *ResponseProfile, error)
	GetProfileByUserID(ctx context.Context, userID string) (*ResponseProfile, error)
	CreateProfile(ctx context.Context, userID string) (*ResponseProfile, error)
}

type RequestUpdateAccount struct {
	UserID      string
	ProfileID   string
	FullName    string
	Gender      string
	Image       *multipart.FileHeader
	PhoneNumber string
	Quote       string
	Profesi     string
}

type ResponseProfile struct {
	ProfileID string
	Quote     *string
	Profesi   *string
}

type ResponseUser struct {
	ID              string
	FullName        string
	Gender          string
	Image           string
	Username        string
	Email           string
	EmailFormat     string
	PhoneNumber     *string
	EmailVerifiedAt bool
}

func (req *RequestUpdateAccount) ToModel(image string) (*repository.Profile, *repository.User) {
	timeUnix := time.Now().Unix()
	profile := &repository.Profile{
		ProfileID: req.ProfileID,
		UserID:    req.UserID,
		Quote:     repository.NewNullString(req.Quote),
		Profesi:   repository.NewNullString(req.Profesi),
		AuditInfo: repository.AuditInfo{
			UpdatedAt: timeUnix,
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}

	user := &repository.User{
		ID:          req.UserID,
		FullName:    req.FullName,
		Gender:      req.Gender,
		Image:       image,
		PhoneNumber: repository.NewNullString(req.PhoneNumber),
		AuditInfo: repository.AuditInfo{
			UpdatedAt: timeUnix,
			UpdatedBy: repository.NewNullString(req.UserID),
		},
	}

	return profile, user
}
