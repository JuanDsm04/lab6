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

// CreateSeries handles the POST request to create a new series in the database.
func CreateSeries(w http.ResponseWriter, r *http.Request) {
	var req Series

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Status == "" {
		respondWithError(w, "Title and status are required", http.StatusBadRequest)
		return
	}

	if req.LastEpisodeWatched < 0 || req.Ranking < 0 {
		respondWithError(w, "Last episode watched and ranking cannot be negative", http.StatusBadRequest)
		return
	}

	if req.TotalEpisodes <= 0 {
		respondWithError(w, "Total episodes must be greater than zero", http.StatusBadRequest)
		return
	}

	if req.LastEpisodeWatched > req.TotalEpisodes {
		respondWithError(w, "Last episode watched cannot be greater than total episodes", http.StatusBadRequest)
		return
	}

	result := db.Exec(
		"INSERT INTO series (title, status, last_episode_watched, total_episodes, ranking) VALUES ($1, $2, $3, $4, $5)",
		req.Title, req.Status, req.LastEpisodeWatched, req.TotalEpisodes, req.Ranking,
	)

	if result.Error != nil {
		log.Println("Error creating series:", result.Error)
		respondWithError(w, "Error creating series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Successful creation",
	})
}

// UpdateSeries handles the PUT request to update an existing series in the database.
func UpdateSeries(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req Series

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Status == "" {
		respondWithError(w, "Title and status are required", http.StatusBadRequest)
		return
	}

	if req.LastEpisodeWatched < 0 || req.Ranking < 0 {
		respondWithError(w, "Last episode watched and ranking cannot be negative", http.StatusBadRequest)
		return
	}

	if req.TotalEpisodes <= 0 {
		respondWithError(w, "Total episodes must be greater than zero", http.StatusBadRequest)
		return
	}

	if req.LastEpisodeWatched > req.TotalEpisodes {
		respondWithError(w, "Last episode watched cannot be greater than total episodes", http.StatusBadRequest)
		return
	}

	result := db.Exec(
		"UPDATE series SET title = ?, status = ?, last_episode_watched = ?, total_episodes = ?, ranking = ? WHERE id = ?",
		req.Title, req.Status, req.LastEpisodeWatched, req.TotalEpisodes, req.Ranking, id,
	)

	if result.Error != nil {
		log.Println("Error updating series:", result.Error)
		respondWithError(w, "Error updating series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Series updated successfully",
		Series:  &req,
	})
}

// DeleteSeries handles the DELETE request to remove a series from the database.
func DeleteSeries(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := db.Delete(&Series{}, id)

	if result.Error != nil {
		respondWithError(w, "Error deleting series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Series deleted",
	})
}

// UpdateSeriesStatus handles the PATCH request to update the status of a series.
func UpdateSeriesStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req struct { Status string `json:"status"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := db.Exec("UPDATE series SET status = ? WHERE id = ?", req.Status, id)
	if result.Error != nil {
		respondWithError(w, "Error updating status", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Status updated",
	})
}
