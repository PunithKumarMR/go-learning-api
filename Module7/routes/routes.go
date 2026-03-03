package routes

import (
	"Module7/handlers"
	"Module7/middleware"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(r *chi.Mux) {
	r.Post("/api/v1/auth/register", handlers.Register)
	r.Post("/api/v1/auth/products/login", handlers.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/api/v1/products", handlers.GetAllProducts)
		r.Get("/api/v1/products/{id}", handlers.GetProduct)
		r.Post("/api/v1/products/{id}", handlers.CreateProduct)
		r.Put("/api/v1/products/{id}", handlers.UpdateProduct)
		r.Delete("/api/v1/products/{id}", handlers.DeleteProduct)
	})
}
