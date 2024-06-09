package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/takumi/coffee-api/controllers"
)

func Route() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/v1/coffees", controllers.GetAllCoffees)
	router.Get("/api/v1/coffee/{id}", controllers.GetCoffeeById)
	router.Post("/api/v1/coffee", controllers.CreateCoffee)
	router.Put("/api/v1/coffee/{id}", controllers.UpdateCoffee)
	router.Delete("/api/v1/coffee/{id}", controllers.DeleteCoffee)

	return router
}
