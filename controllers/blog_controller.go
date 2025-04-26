package controllers

import (
	"blog-crud-api/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Simulated DB
var blogPosts []models.BlogPost
var idCounter = 1

// CreateBlogPost godoc
// @Summary Create a blog post
// @Produce json
// @Param blog body models.BlogPost true "Blog Post"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post [post]
func CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var blog models.BlogPost
	_ = json.NewDecoder(r.Body).Decode(&blog)

	blog.ID = idCounter
	blog.CreatedAt = time.Now().Format(time.RFC3339)
	blog.UpdatedAt = blog.CreatedAt

	idCounter++
	blogPosts = append(blogPosts, blog)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

// GetAllBlogPosts godoc
// @Summary Get all blog posts
// @Produce json
// @Success 200 {array} models.BlogPost
// @Router /api/blog-post [get]
func GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogPosts)
}

// GetBlogPost godoc
// @Summary Get a blog post by ID
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [get]
func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range blogPosts {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// UpdateBlogPost godoc
// @Summary Update a blog post
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body models.BlogPost true "Blog Post"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [patch]
func UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range blogPosts {
		if item.ID == id {
			var updatedBlog models.BlogPost
			_ = json.NewDecoder(r.Body).Decode(&updatedBlog)

			updatedBlog.ID = id
			updatedBlog.CreatedAt = item.CreatedAt
			updatedBlog.UpdatedAt = time.Now().Format(time.RFC3339)

			blogPosts[index] = updatedBlog

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBlog)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// DeleteBlogPost godoc
// @Summary Delete a blog post
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {string} string "Deleted Successfully"
// @Router /api/blog-post/{id} [delete]
func DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range blogPosts {
		if item.ID == id {
			blogPosts = append(blogPosts[:index], blogPosts[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}
