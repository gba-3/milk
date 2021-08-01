package repository

import (
	"fmt"
	"log"

	"github.com/gba-3/milk/domain/entity"
	"github.com/gba-3/milk/infrastructure"
	"github.com/gba-3/milk/logger"
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
	var users []entity.User
	query := "SELECT * FROM users"
	err := infrastructure.NewMySQL(ur.db).Select(&users, query)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func (ur *userRepository) CreateUser(name string, email string, password string) error {
	logger.Log.Info(fmt.Sprintf("name: %s, email: %s", name, email))
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	return infrastructure.NewMySQL(ur.db).Exec(query, name, email, password)
}
