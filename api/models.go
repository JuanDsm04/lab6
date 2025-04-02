// This file contains the definition of the Series and ApiResponse structs.
package main

// Series represents a TV series or Anime with its attributes.
type Series struct {
    ID                 uint   `gorm:"column:id;primaryKey" json:"id"`
    Title              string `gorm:"column:title" json:"title"`
    Status             string `gorm:"column:status" json:"status"`
    LastEpisodeWatched int    `gorm:"column:last_episode_watched" json:"lastEpisodeWatched"`
    TotalEpisodes      int    `gorm:"column:total_episodes" json:"totalEpisodes"`
    Ranking            int    `gorm:"column:ranking" json:"ranking"`
}

// ApiResponse represents the structure of an API response.
type ApiResponse struct {
	Success bool    `json:"success"`	// Indicates whether the request was successful
	Message string  `json:"message"`	// Message related to the response
	Series  *Series `json:"series,omitempty"`	// Optional series data, included if available
}