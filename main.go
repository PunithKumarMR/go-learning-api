package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
func getIDFromURL(r *http.Request) (int, error) {
	idStr := r.URL.Path[len("/api/v1/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, err
}
func allProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(products)
	case http.MethodPost:
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
	default:
		http.Error(w, "Mehtod Not Allowd", http.StatusMethodNotAllowed)
	}
}
func singleProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := getIDFromURL(r)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	index, product := findProductByID(id)
	if product == nil {
		http.Error(w, "Product Not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(product)
	case http.MethodPut:
		var updated Product
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated.ID = id
		products[index] = updated
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		products = append(products[:index], products[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Metheod Not Allowed", http.StatusMethodNotAllowed)
	}

}

func RequestHandler() {
	http.HandleFunc("/api/v1/products", allProductsHandler)
	http.HandleFunc("/api/v1/products/", singleProductHandler)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func main() {
	RequestHandler()
}
