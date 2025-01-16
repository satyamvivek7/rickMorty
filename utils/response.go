package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends a JSON response to the client
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

