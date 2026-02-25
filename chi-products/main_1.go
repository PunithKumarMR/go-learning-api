package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []Product{}

func findProductByID(id int) (int, *Product) {
	for i, p := range products {
		if p.ID == id {
			return i, &p
		}
	}
	return -1, nil
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, product := findProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	index, product := findProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	var updated Product
	err1 := json.NewDecoder(r.Body).Decode(&updated)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	updated.ID = id
	products[index] = updated
	json.NewEncoder(w).Encode(updated)
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	index, product := findProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	products = append(products[:index], products[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	r := chi.NewRouter()
	r.Get("/api/v1/products", getAllProducts)
	r.Post("/api/v1/products", createProduct)
	r.Get("/api/v1/products/{id}", getProduct)
	r.Put("/api/v1/products/{id}", updateProduct)
	r.Delete("/api/v1/products/{id}", deleteProduct)
	log.Fatal(http.ListenAndServe(":8085", r))

}
