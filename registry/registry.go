package registry

import (
	"github.com/gba-3/milk/registry/container"
)

type registry struct {
	c container.Container
}

func NewRegistry() *registry {
	return &registry{
		c: container.Container{},
	}
}

func (r *registry) GetAppHandler() *AppHandler {
	handler := NewAppHandler(
		r.c.GetUserHandler(
			r.c.GetUserUsecase(
				r.c.GetUserRepository(),
			),
		),
	)
	return handler
}
