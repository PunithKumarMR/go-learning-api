package routes

import (
	"Module4/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Get("/api/v1/products", handlers.GetAllProducts)
	r.Get("/api/v1/products/{id}", handlers.GetProduct)
	r.Post("/api/v1/products", handlers.CreateProduct)
	r.Put("/api/v1/products/{id}", handlers.UpdateProduct)
	r.Delete("/api/v1/products/{id}", handlers.DeleteProduct)
}
