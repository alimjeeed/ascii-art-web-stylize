package main

import (
	"asciiart/web"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes and their handlers.
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/ascii-art", web.AsciiArtHandler)

	log.Println("Server started on: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
