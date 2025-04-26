package routes

import (
	"blog-crud-api/controllers"

	"github.com/gorilla/mux"
)

func BlogRoutes(router *mux.Router) {
	router.HandleFunc("/api/blog-post", controllers.CreateBlogPost).Methods("POST")
	router.HandleFunc("/api/blog-post", controllers.GetAllBlogPosts).Methods("GET")
	router.HandleFunc("/api/blog-post/{id}", controllers.GetBlogPost).Methods("GET")
	router.HandleFunc("/api/blog-post/{id}", controllers.UpdateBlogPost).Methods("PATCH")
	router.HandleFunc("/api/blog-post/{id}", controllers.DeleteBlogPost).Methods("DELETE")
}
