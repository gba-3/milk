package handler

import (
	"net/http"

	"github.com/gba-3/milk/usecase"
)

type UserHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{uu}
}

func (uh UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	users := uh.uu.GetUsers()
	return http.StatusOK, users, nil
}
