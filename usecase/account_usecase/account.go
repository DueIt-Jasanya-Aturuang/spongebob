package account_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/usecase"
)

type AccountUsecaseImpl struct {
	profileRepo repository.ProfileRepository
	userRepo    repository.UserRepository
	minioRepo   repository.MinioRepository
}

func NewAccountUsecaseImpl(
	profileRepo repository.ProfileRepository,
	userRepo repository.UserRepository,
	minioRepo repository.MinioRepository,
) usecase.AccountUsecase {
	return &AccountUsecaseImpl{
		profileRepo: profileRepo,
		userRepo:    userRepo,
		minioRepo:   minioRepo,
	}
}
