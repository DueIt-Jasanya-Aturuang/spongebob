package user_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
)

type UserRepoImpl struct {
	repository.UnitOfWorkRepository
}

func NewUserRepositoryImpl(uow repository.UnitOfWorkRepository) repository.UserRepository {
	return &UserRepoImpl{
		UnitOfWorkRepository: uow,
	}
}
