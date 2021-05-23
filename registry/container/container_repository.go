package container

import "github.com/gba-3/milk/domain/repository"

func (c Container) GetUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
