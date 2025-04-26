package main

import (
	"blog-crud-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "blog-crud-api/docs" // Important for swagger
)

// @title Blog CRUD API
// @version 1.0
// @description This is a simple blog CRUD API in Go.
// @host localhost:8080
// @BasePath /

func main() {
	router := mux.NewRouter()

	routes.BlogRoutes(router)

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
