package profileConfig_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type ProfileConfigRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewProfileConfigRepositoryImpl(uow repository.UnitOfWorkRepository) repository.ProfileConfigRepository {
	return &ProfileConfigRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
