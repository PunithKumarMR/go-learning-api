package main

import (
	"Module4/database"
	"Module4/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()
	route := chi.NewRouter()
	routes.SetupRoutes(route)
	log.Fatal(http.ListenAndServe(":8086", route))

}
