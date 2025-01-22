package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"orem/config"
	"orem/handlers"
	"orem/handlers/middleware"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	// Configure structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Initialize configuration
	config.Initialize()

	// Setup routes with logging middleware
	http.HandleFunc("GET /activities", middleware.WithLogging(handlers.GetAllActivities))
	http.HandleFunc("GET /activities/{id}", middleware.WithLogging(handlers.GetActivityDetails))

	cfg := config.Get()
	slog.Info("Server is running", "port", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
