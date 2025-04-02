// This file contains functions to send JSON responses for success or error.

package main

import (
	"encoding/json"
	"net/http"
)

// respondWithJSON sends a JSON response to the client.
// w: the HTTP response writer.
// payload: the data to send as a JSON response.
func respondWithJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

// respondWithError sends an error response in JSON format.
// w: the HTTP response writer.
// message: the error message to send.
// status: the HTTP status code.
func respondWithError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Message: message,
	})
}