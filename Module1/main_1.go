package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var groceries = []Todo{}

func findItemByID(id int) (int, *Todo) {
	for i, item := range groceries {
		if item.ID == id {
			return i, &item
		}
	}
	return -1, nil
}
func findIDByURL(r *http.Request) (int, error) {
	idStr := r.URL.Path[len("/api/v1/todolist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, err
}

func singleItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := findIDByURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	index, item := findItemByID(id)
	if item == nil {
		http.Error(w, "Item Not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(item)
	case http.MethodPut:
		var updated Todo
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated.ID = id
		groceries[index] = updated
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		groceries = append(groceries[:index], groceries[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	}

}

func allItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(groceries)
	case http.MethodPost:
		var newItem Todo

		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newItem.ID = len(groceries) + 1
		groceries = append(groceries, newItem)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
	default:
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	}

}
func RequestHandler() {
	http.HandleFunc("/api/v1/todolist/", singleItemHandler)
	http.HandleFunc("/api/v1/todolist", allItemHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func main() {
	RequestHandler()
}
