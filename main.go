package main

import (
	"log"
	"net/http"
	"rickmorty/config"
	"rickmorty/database"
	"rickmorty/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configurations
	config.LoadEnv()

	// Connect to MongoDB
	client, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/character", handlers.GetCharacter).Methods("GET")

	// Start the server
	port := "8080"
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}


