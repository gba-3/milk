package usecase

import (
	"testing"

	"github.com/gba-3/milk/domain/entity"
)

type mockUserRepository struct{}

func (mur *mockUserRepository) GetUsers() []entity.User {
	users := []entity.User{
		{
			ID:       1,
			Name:     "user1",
			Email:    "user1@email.com",
			Password: "testpass",
		},
		{
			ID:       2,
			Name:     "user2",
			Email:    "user2@email.com",
			Password: "testpass",
		},
		{
			ID:       3,
			Name:     "user3",
			Email:    "user3@email.com",
			Password: "testpass",
		},
	}
	return users
}

type mockEmptyUserRepository struct{}

func (meur mockEmptyUserRepository) GetUsers() []entity.User {
	return []entity.User{}
}

func TestSuccessGetUsers(t *testing.T) {
	ur := &mockUserRepository{}
	uu := NewUserUsecase(ur)
	users := uu.GetUsers()
	if len(users) == 0 {
		t.Fatal("Users table is empty.")
	}
}

func TestEmptyGetUsers(t *testing.T) {
	ur := &mockEmptyUserRepository{}
	uu := NewUserUsecase(ur)
	users := uu.GetUsers()
	if len(users) > 0 {
		t.Fatal("Users table is not empty.")
	}
}
