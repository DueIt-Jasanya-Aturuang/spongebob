package profileConfig_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

type ProfileConfigUsecaseImpl struct {
	profileRepo    repository.ProfileRepository
	profileCfgRepo repository.ProfileConfigRepository
	notifRepo      repository.NotificationRepository
}

func NewProfileConfigUsecaseImpl(
	profileRepo repository.ProfileRepository,
	profileCfgRepo repository.ProfileConfigRepository,
	notifRepo repository.NotificationRepository,
) usecase.ProfileConfigUsecase {
	return &ProfileConfigUsecaseImpl{
		profileRepo:    profileRepo,
		profileCfgRepo: profileCfgRepo,
		notifRepo:      notifRepo,
	}
}
