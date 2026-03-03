package routes

import (
	"Module6/handlers"
	"Module6/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	// public routes — no auth needed
	r.Post("/api/v1/register", handlers.Register)
	r.Post("/api/v1/login", handlers.Login)

	// protected routes — auth needed
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/api/v1/products", handlers.GetAllProducts)
		r.Get("/api/v1/products/{id}", handlers.GetProduct)
		r.Post("/api/v1/products", handlers.CreateProduct)
		r.Put("/api/v1/products/{id}", handlers.UpdateProduct)
		r.Delete("/api/v1/products/{id}", handlers.DeleteProduct)
	})
}
