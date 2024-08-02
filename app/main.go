package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define the port to listen on
	port := "8080"

	// Check if a port is specified in the environment variables
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Create a file server handler
	fs := http.FileServer(http.Dir("."))

	// Set up the handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the index.html file
		http.ServeFile(w, r, "index.html")
	})

	// Set up a handler for all other paths to serve static files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
