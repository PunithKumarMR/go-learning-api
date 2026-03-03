package main

import (
	"Module6/database"
	"Module6/routes"
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
