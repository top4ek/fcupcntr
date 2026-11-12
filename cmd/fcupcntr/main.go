package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fcupcntr/internal/api"
	"fcupcntr/internal/config"
	"fcupcntr/internal/sqlite"
)

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := config.Load()
	if err != nil {
		return err
	}

	database, err := sqlite.Open(ctx, config.Settings.DSN)
	if err != nil {
		return err
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Printf("Database closing error: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", api.HealthHandler)

	server := &http.Server{
		Addr:         ":" + config.Settings.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	serverErrors := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case <-ctx.Done():
	case err := <-serverErrors:
		return err
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}
	log.Println("Server stopped.")
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Start error: %v", err)
	}
}
