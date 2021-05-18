package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/api", func(api chi.Router) {
		api.Route("/signup", func(signup chi.Router) {
			signup.Post("/", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("Hello, World"))
			})
		})
	})

	http.ListenAndServe(":9090", router)
}
