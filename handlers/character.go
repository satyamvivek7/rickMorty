// package handlers

// import (
// 	// "encoding/json"
// 	"net/http"
// 	"rickmorty/models"
// 	"rickmorty/utils"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// // GetCharacter handles the API for searching characters by name
// func GetCharacter(w http.ResponseWriter, r *http.Request) {
// 	// Get the character name from query parameters
// 	name := r.URL.Query().Get("name")
// 	if name == "" {
// 		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "name query parameter is required"})
// 		return
// 	}

// 	// Access the database and collection
// 	collection := client.Database("RickMorty").Collection("characters")

// 	// Search for characters by name
// 	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
// 	cursor, err := collection.Find(r.Context(), filter, options.Find())
// 	if err != nil {
// 		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
// 		return
// 	}
// 	defer cursor.Close(r.Context())

// 	// Parse results
// 	var characters []models.Character
// 	if err := cursor.All(r.Context(), &characters); err != nil {
// 		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to parse results"})
// 		return
// 	}

// 	utils.RespondWithJSON(w, http.StatusOK, characters)
// }


//---------------------------------------------------------------//
// 


package handlers

import (
	"log"
	"net/http"
	"rickmorty/models"
	"rickmorty/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// GetCharacter handles the API for searching characters by name
func GetCharacter(w http.ResponseWriter, r *http.Request) {
	// Ensure MongoDB client is initialized
	if client == nil {
		log.Println("MongoDB client is not initialized")
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "MongoDB client is not initialized"})
		return
	}

	// Get the character name from query parameters
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "name query parameter is required"})
		return
	}

	// Access the database and collection
	collection := client.Database("RickMorty").Collection("characters")

	// Search for characters by name using a case-insensitive regex
	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
	cursor, err := collection.Find(r.Context(), filter, options.Find())
	if err != nil {
		log.Println("Database query failed:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		return
	}
	defer cursor.Close(r.Context())

	// Parse results into characters slice
	var characters []models.Character
	if err := cursor.All(r.Context(), &characters); err != nil {
		log.Println("Failed to parse results:", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to parse results"})
		return
	}

	// Respond with the characters in JSON format
	utils.RespondWithJSON(w, http.StatusOK, characters)
}
