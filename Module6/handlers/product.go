package handlers

import (
	"Module6/database"
	"Module6/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var product models.Product
	find := database.DB.First(&product, id)
	if find.Error != nil {
		http.Error(w, "Product not found!", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	err1 := json.NewDecoder(r.Body).Decode(&newProduct)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	create := database.DB.Create(&newProduct)
	if create.Error != nil {
		http.Error(w, create.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var Updated models.Product
	find := database.DB.First(&Updated, id)
	if find.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	err1 := json.NewDecoder(r.Body).Decode(&Updated)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	save := database.DB.Save(&Updated)
	if save.Error != nil {
		http.Error(w, save.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Updated)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var Updated models.Product
	find := database.DB.First(&Updated, id)
	if find.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	result := database.DB.Delete(&Updated)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
