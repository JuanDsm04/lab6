package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
)

// GetAllSeries retrieves all series.
// @Summary Retrieves all series
// @Description Returns a list of series stored in the database.
// @Tags series
// @Produce json
// @Success 200 {array} Series
// @Router /api/series [get]
func GetAllSeries(w http.ResponseWriter, r *http.Request) {
	var series []Series
	result := db.Find(&series)

	if result.Error != nil {
		respondWithError(w, "Error retrieving series", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, series)
}

// GetSeriesByID retrieves a specific series by its ID.
// @Summary Retrieves a series by ID
// @Description Returns a series stored in the database based on the provided ID.
// @Tags series
// @Produce json
// @Param id path int true "Series ID"
// @Success 200 {object} Series
// @Router /api/series/{id} [get]
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

// CreateSeries adds a new series.
// @Summary Creates a new series
// @Description Allows adding a series to the database.
// @Tags series
// @Accept json
// @Produce json
// @Param series body Series true "Series data"
// @Success 201 {object} Series
// @Router /api/series [post]
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

// UpdateSeries updates an existing series.
// @Summary Updates a series
// @Description Allows modifying a series in the database.
// @Tags series
// @Accept json
// @Produce json
// @Param id path int true "Series ID"
// @Param series body Series true "Updated series data"
// @Success 200 {object} Series
// @Router /api/series/{id} [put]
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

// DeleteSeries removes a series.
// @Summary Deletes a series
// @Description Deletes a series based on its ID.
// @Tags series
// @Param id path int true "Series ID"
// @Success 200 {string} string "Series deleted"
// @Router /api/series/{id} [delete]
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

// UpdateSeriesStatus updates the status of a series.
// @Summary Update the status of a series
// @Description Updates the status field of a series in the database.
// @Tags series
// @Accept json
// @Produce json
// @Param id path int true "Series ID"
// @Param status body string true "New status"
// @Success 200 {object} ApiResponse
// @Router /api/series/{id}/status [patch]
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

// IncrementEpisode increments the last episode watched of a series.
// @Summary Increments the last episode watched
// @Description Increments the last_episode_watched field of a series.
// @Tags series
// @Param id path int true "Series ID"
// @Success 200 {object} ApiResponse
// @Router /api/series/{id}/episode [patch]
func IncrementEpisode(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := db.Exec("UPDATE series SET last_episode_watched = last_episode_watched + 1 WHERE id = ?", id)
	if result.Error != nil {
		respondWithError(w, "Error incrementing episode", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Episode incremented",
	})
}

// UpvoteSeries increases the ranking of a series.
// @Summary Upvote a series
// @Description Increases the ranking field of a series by 1.
// @Tags series
// @Param id path int true "Series ID"
// @Success 200 {object} ApiResponse
// @Router /api/series/{id}/upvote [patch]
func UpvoteSeries(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := db.Exec("UPDATE series SET ranking = ranking + 1 WHERE id = ?", id)
	if result.Error != nil {
		respondWithError(w, "Error upvoting series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Series upvoted",
	})
}

// DownvoteSeries decreases the ranking of a series by 1.
// @Summary Downvotes a series
// @Description Decreases the ranking field of a series by 1.
// @Tags series
// @Param id path int true "Series ID"
// @Success 200 {object} ApiResponse
// @Router /api/series/{id}/downvote [patch]
func DownvoteSeries(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result := db.Exec("UPDATE series SET ranking = ranking - 1 WHERE id = ?", id)
	if result.Error != nil {
		respondWithError(w, "Error downvoting series", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Series downvoted",
	})
}