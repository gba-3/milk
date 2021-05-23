package repository

import "github.com/gba-3/milk/domain/entity"

type userRepository struct {
}

type UserRepository interface {
	GetUsers() []entity.User
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUsers() []entity.User {
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
