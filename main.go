package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jjlock/byte-scraper-api/handler"
)

func main() {
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      handler.NewScraperHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	serverClosed := make(chan struct{})
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server ListenAndServe: %v", err)
		}
		close(serverClosed)
	}()
	log.Println("Server started on localhost" + srv.Addr)

	<-done
	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// defers are not called with os.Exit so call cancel here
		cancel()
		log.Printf("Server shutdown failed: %v", err)
	}

	<-serverClosed
	log.Println("Server shutdown gracefully")
}
