package main

import (
	"encoding/json"
	"net/http"
)

// GetAllSeries handles the GET request to retrieve all series from the database.
func GetAllSeries(w http.ResponseWriter, r *http.Request) {
	var series []Series
	result := db.Find(&series)

	if result.Error != nil {
		respondWithError(w, "Error retrieving series", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, series)
}