package main

import (
	"github.com/gba-3/milk/auth"
	"log"
	"net/http"

	"github.com/gba-3/milk/handler"
	"github.com/gba-3/milk/infrastructure"
	"github.com/gba-3/milk/registry"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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
			signup.Post("/", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("Signup"))
			})
		})
		api.Route("/users", func(users chi.Router) {
			users.Get("/", auth.JwtMiddleware.Handler(
				handler.JsonHandler(ah.UserHandler.GetUsers)).ServeHTTP,
			)
			users.Post("/signup", handler.JsonHandler(
				ah.UserHandler.Signup,
			).ServeHTTP)
		})
	})

	http.ListenAndServe(":9090", router)
}
