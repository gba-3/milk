package registry

import (
	"github.com/gba-3/milk/handler"
)

type AppHandler struct {
	UserHandler *handler.UserHandler
}

func NewAppHandler(uh *handler.UserHandler) *AppHandler {
	return &AppHandler{uh}
}
