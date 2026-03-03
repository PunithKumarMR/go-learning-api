package main

import (
	"Module7/database"
	"Module7/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()
	r := chi.NewRouter()
	routes.SetUpRoutes(r)
	log.Fatal(http.ListenAndServe(":8082", r))
}
