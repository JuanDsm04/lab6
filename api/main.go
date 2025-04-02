package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Define the data source name (DSN) for connecting to the PostgreSQL database.
	dsn := "host=db user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable TimeZone=UTC"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	// Initialize a new router using the chi package.
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/api/series", GetAllSeries)
	r.Get("/api/series/{id}", GetSeriesByID)
	r.Post("/api/series", CreateSeries)
	r.Put("/api/series/{id}", UpdateSeries)
	r.Delete("/api/series/{id}", DeleteSeries)

	r.Patch("/api/series/{id}/status", UpdateSeriesStatus)
	r.Patch("/api/series/{id}/episode", IncrementEpisode)
	r.Patch("/api/series/{id}/upvote", UpvoteSeries)

	// Start the server and log the URL.
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}