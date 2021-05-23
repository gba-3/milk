package usecase

import (
	"github.com/gba-3/milk/domain/entity"
	"github.com/gba-3/milk/domain/repository"
)

type userUsecase struct {
	ur repository.UserRepository
}

type UserUsecase interface {
	GetUsers() []entity.User
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) GetUsers() []entity.User {
	return uu.ur.GetUsers()
}
