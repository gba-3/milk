package repository

import (
	"log"

	"github.com/gba-3/milk/domain/entity"
	"github.com/gba-3/milk/infrastructure"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	GetUsers() []entity.User
	CreateUser(name string, email string, password string) error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUsers() []entity.User {
	// users := []entity.User{
	// 	{
	// 		ID:       1,
	// 		Name:     "user1",
	// 		Email:    "user1@email.com",
	// 		Password: "testpass",
	// 	},
	// 	{
	// 		ID:       2,
	// 		Name:     "user2",
	// 		Email:    "user2@email.com",
	// 		Password: "testpass",
	// 	},
	// 	{
	// 		ID:       3,
	// 		Name:     "user3",
	// 		Email:    "user3@email.com",
	// 		Password: "testpass",
	// 	},
	// }
	var users []entity.User
	query := "SELECT * FROM users"
	err := infrastructure.NewMySQL(ur.db).Select(&users, query)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func (ur *userRepository) CreateUser(name string, email string, password string) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	return infrastructure.NewMySQL(ur.db).Exec(query, name, email, password)
}
