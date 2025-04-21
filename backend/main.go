package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"utilss"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	// Loading .env variables (sensible informations)
	utils.LoadEnv(".env")

	// Database / Websocket hub setup

	// Router configuration
	mux := setupMux()
	// Need rate limiter implementation
	server := setupServer(mux)

	// Configuration HTTPS
	// lib.SetupHTTPS(server)

	log.Printf("Server starting on https://%s...\n", server.Addr)

	// Démarrage du serveur HTTPS
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil { // Utilisation de ListenAndServeTLS pour HTTPS
		log.Fatalf("Error starting server: %v", err)
	}
}

// Configuration du routeur
func setupMux() *http.ServeMux {
	// New ServeMux setup
	mux := http.NewServeMux()

	// Handling static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/static/uploads/", http.StripPrefix("/static/uploads/", http.FileServer(http.Dir("./static/uploads/"))))

	// Set up routes
	mux.HandleFunc("/", handlers.HomeHandler)

	// Basic Web handlers
	// mux.HandleFunc("/about", handlers.AboutHandler)
	//mux.HandleFunc("/error", handlers.ForceDirectError) // !for testing purpose only (not for production)
	//mux.HandleFunc("/500", handlers.Force500Handler)    // !for testing purpose only (not for production)

	return mux
}

// Configuration du serveur
func setupServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              ":8080",
		Handler:           utils.WithErrorHandling(handler),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
