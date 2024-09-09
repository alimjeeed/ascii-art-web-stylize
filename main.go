package main

import (
	"asciiart/web"
	"log"
	"net/http"
)

func main() {
	// Set up routes and their handlers.
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/ascii-art", web.AsciiArtHandler)

	// Log server start message.
	log.Println("Server started on: http://localhost:8000")

	// Start the HTTP server on port 8000.
	// log.Fatal ensures the server stops if an error occurs.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
