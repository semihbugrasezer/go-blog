package main

import (
	"blog-platform/controllers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:0606", "http://localhost:5500", "http://127.0.0.1:5500/"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // OPTIONS method for CORS preflight
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// Apply the CORS middleware to the router
	handler := c.Handler(router)

	// Start the server and log any errors
	log.Println("Starting server on :0606")
	if err := http.ListenAndServe(":0606", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

}
