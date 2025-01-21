package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"orem/handlers"
	"orem/handlers/middleware"
	"os"
)

func main() {
	// Configure structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Setup routes with logging middleware
	http.HandleFunc("GET /activities", middleware.WithLogging(handlers.GetAllActivities))
	http.HandleFunc("GET /activities/{id}", middleware.WithLogging(handlers.GetActivityDetails))

	port := 8628
	slog.Info("Server is running", "port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
