package registry

import (
	"github.com/gba-3/milk/registry/container"
	"github.com/jmoiron/sqlx"
)

type registry struct {
	c container.Container
}

func NewRegistry() *registry {
	return &registry{
		c: container.Container{},
	}
}

func (r *registry) GetAppHandler(db *sqlx.DB) *AppHandler {
	handler := NewAppHandler(
		r.c.GetUserHandler(
			r.c.GetUserUsecase(
				r.c.GetUserRepository(db),
			),
		),
	)
	return handler
}
