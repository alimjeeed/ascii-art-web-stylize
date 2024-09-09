package main

import (
	"asciiart/web"
	"log"
	"net/http"
)

func main() {
	// Solves external .CSS file not working.
	// Serve static files from the "static" directory.
	// Strip the "/static/" prefix from the URL path before forwarding the request to the file server.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes and their handlers.
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/ascii-art", web.AsciiArtHandler)

	// Log server start message.
	log.Println("Server started on: http://localhost:8000")

	// Start the HTTP server on port 8000.
	// log.Fatal ensures the server stops if an error occurs.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
