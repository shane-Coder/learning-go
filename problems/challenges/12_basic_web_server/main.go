package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home page!")
}

func aboutHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "This is the About page!")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Contact us at support@example.com")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on port :8080...")
	fmt.Println("Visit http://localhost:8080 in your browser.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
