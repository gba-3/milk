package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gba-3/milk/domain/entity"
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

func (uh UserHandler) Signup(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var reqBody entity.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if json.Unmarshal(body, &reqBody); err != nil {
	}
}
