package handler

import (
	"encoding/json"
	"errors"
	"github.com/gba-3/milk/auth"
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
		return http.StatusBadRequest, nil, err
	}
	if reqBody.Name == "" {
		return http.StatusBadRequest, nil, errors.New("Name is empty in request body.")
	}
	if reqBody.Email == "" {
		return http.StatusBadRequest, nil, errors.New("Email is empty in request body.")
	}
	if reqBody.Password == "" {
		return http.StatusBadRequest, nil, errors.New("Password is empty in request body.")
	}
	err = uh.uu.CreateUser(reqBody.Name, reqBody.Email, reqBody.Password)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	token, err := auth.CreateToken(reqBody.Email)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	res := map[string]string{
		"token": token,
	}
	return http.StatusOK, res, nil
}
