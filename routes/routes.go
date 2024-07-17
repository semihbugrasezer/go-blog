package routes

import (
	"blog-platform/controllers"
	"blog-platform/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	api.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	api.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")
	api.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET")

	return r
}
