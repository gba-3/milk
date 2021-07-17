package container

import (
	"github.com/gba-3/milk/domain/repository"
	"github.com/jmoiron/sqlx"
)

func (c Container) GetUserRepository(db *sqlx.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}
