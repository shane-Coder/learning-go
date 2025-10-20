package main

import (
	"fmt"
	"log" // Import the log package for better error handling
	"net/http"
)

func main() {
	// 1. Create the file server handler.
	// http.Dir specifies the directory to serve files from.
	// http.FileServer creates the handler.
	fileServer := http.FileServer(http.Dir("./public"))

	// 2. Register the file server handler for the root path "/".
	// We use http.Handle because fileServer is already a handler object.
	http.Handle("/", fileServer)

	// 3. Print startup messages.
	fmt.Println("Starting server on port :8080...")
	fmt.Println("Visit http://localhost:8080 in your browser.")

	// 4. Start the server and handle potential errors.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err) // Use log.Fatalf for critical errors
	}
}
