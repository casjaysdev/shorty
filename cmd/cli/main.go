// File: cmd/server/main.go
// Purpose: Entry point for the Shorty server. Loads config, routes, logging, and runs the HTTP server.

package main

import (
	"log"
	"net/http"
	"os"

	"shorty/internal/config"
	"shorty/internal/lib/logger"
	"shorty/internal/lib/middleware"
	"shorty/internal/router"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup logging
	log := logger.New(cfg.LogLevel)

	// Initialize router
	r := router.Setup(cfg)

	// Attach global middleware
	handler := middleware.WithGlobal(r, cfg)

	// Run HTTP server
	addr := cfg.Server.ListenAddr
	log.Infof("Starting Shorty server on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server error: %v", err)
		os.Exit(1)
	}
}
