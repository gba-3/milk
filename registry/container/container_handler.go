package container

import (
	"github.com/gba-3/milk/handler"
	"github.com/gba-3/milk/usecase"
)

func (c Container) GetUserHandler(uu usecase.UserUsecase) *handler.UserHandler {
	return handler.NewUserHandler(uu)
}
