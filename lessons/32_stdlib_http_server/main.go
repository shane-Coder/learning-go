package main

import (
	"fmt"
	"net/http"
)

// handlerForRoot is a handler for the "/" path.
func handlerForRoot(w http.ResponseWriter, r *http.Request) {
	// Fprintf writes a formatted string to a Writer.
	// The http.ResponseWriter 'w' is a Writer!
	fmt.Fprintf(w, "Hello, you've reached the home page!")
}

// handlerForAbout is a handler for the "/about" path.
func handlerForAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page.")
}

func main() {
	// Register our handler functions for specific URL paths.
	http.HandleFunc("/", handlerForRoot)
	http.HandleFunc("/about", handlerForAbout)

	// Print a message to the console to show the server is starting.
	fmt.Println("Starting server on port :8080...")
	fmt.Println("Visit http://localhost:8080 in your browser.")

	// http.ListenAndServe starts the server.
	// It takes the port to listen on and a "handler". We pass nil to use
	// the default handler we just configured.
	// This function is blocking; it will run forever unless an error occurs.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
