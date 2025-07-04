// File: cmd/server/main.go
// Purpose: Entry point of the shorty backend server; wires everything together and starts the HTTP service.

package main

import (
	"context"
	"os"
	"os/signal"
	"shorty/internal/config"
	"shorty/internal/lib/logger"
	"shorty/internal/server"
	"shorty/internal/storage"
	"shorty/internal/version"
	"syscall"
	"time"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.New(cfg)

	// Print version info
	log.Infof("Starting Shorty server - Version: %s Commit: %s Date: %s", version.BuildVersion, version.BuildCommit, version.BuildDate)

	// Initialize database
	db, err := storage.Init(cfg, log)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize HTTP server
	srv := server.New(cfg, log, db)

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Infof("Shutting down...")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxTimeout); err != nil {
		log.Errorf("Shutdown error: %v", err)
	}

	log.Infof("Shutdown complete")
}
