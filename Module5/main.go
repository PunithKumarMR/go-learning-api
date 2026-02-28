package main

import (
	"Module5/database"
	"Module5/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()
	r := chi.NewRouter()
	routes.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(":8087", r))
}
