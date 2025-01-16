package models

// Character represents a Rick and Morty character
type Character struct {
	ID       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Status   string   `json:"status" bson:"status"`
	Species  string   `json:"species" bson:"species"`
	Gender   string   `json:"gender" bson:"gender"`
	Episodes []string `json:"episodes" bson:"episodes"`
}

