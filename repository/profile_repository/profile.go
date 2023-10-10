package profile_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type ProfileRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewProfileRepositoryImpl(uow repository.UnitOfWorkRepository) repository.ProfileRepository {
	return &ProfileRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
