package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gba-3/milk/auth"
	"github.com/gba-3/milk/logger"

	"github.com/gba-3/milk/handler"
	"github.com/gba-3/milk/infrastructure"
	"github.com/gba-3/milk/registry"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func init() {
	level := os.Getenv("LOGLEVEL")
	logger.SetupLogger(level)
}

func main() {
	db, err := infrastructure.BootMySQL()
	if err != nil {
		log.Fatal(err)
	}
	r := registry.NewRegistry()
	ah := r.GetAppHandler(db)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/api", func(api chi.Router) {
		api.Route("/signup", func(signup chi.Router) {
			signup.Post("/", handler.JsonHandler(
				ah.UserHandler.Signup,
			).ServeHTTP)
		})
		api.Route("/users", func(users chi.Router) {
			users.Get("/", auth.JwtMiddleware.Handler(
				handler.JsonHandler(ah.UserHandler.GetUsers)).ServeHTTP,
			)
		})
	})

	if err = http.ListenAndServe(":9090", router); err != nil {
		logger.Log.Error(err.Error())
	}
}
