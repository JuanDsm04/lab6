package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
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

// GetSeriesByID handles the GET request to retrieve a specific series by its ID from the database.
func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var series Series
	result := db.First(&series, id)

	if result.Error != nil {
		respondWithError(w, "Series not found", http.StatusNotFound)
		return
	}
	respondWithJSON(w, series)
}