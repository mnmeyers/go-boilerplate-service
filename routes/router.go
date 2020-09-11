package routes

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"go-boilerplate/controllers"
	"net/http"
)

// GetRouter returns the mux that handles all incoming HTTP Requests
func GetRouter() http.Handler {
	router, c := makeRouter()

	return c.Handler(router)
}

func makeRouter() (*chi.Mux, *cors.Cors) {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "POST", "post", "DELETE", "PATCH", "patch"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
		Debug:            false,
	})

	router.Use(c.Handler)

	router.Route("/", func(innerRouter chi.Router) {
		controller := controllers.GetController()

		innerRouter.Post("/things", controller.Post)
		innerRouter.Get("/{account_id}/things", controller.List)
	})

	return router, c
}
