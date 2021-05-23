package main

import (
	"net/http"

	"github.com/gba-3/milk/handler"
	"github.com/gba-3/milk/registry"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := registry.NewRegistry()
	ah := r.GetAppHandler()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/api", func(api chi.Router) {
		api.Route("/signup", func(signup chi.Router) {
			signup.Post("/", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("Signup"))
			})
		})
		api.Route("/users", func(users chi.Router) {
			users.Get("/", handler.JsonHandler(
				ah.UserHandler.GetUsers,
			).ServeHTTP)
		})
	})

	http.ListenAndServe(":9090", router)
}
