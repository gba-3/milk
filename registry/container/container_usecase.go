package container

import (
	"github.com/gba-3/milk/domain/repository"
	"github.com/gba-3/milk/usecase"
)

func (c Container) GetUserUsecase(ur repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(ur)
}
