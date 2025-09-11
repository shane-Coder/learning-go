/*
Problem Statement #16: Simple JSON API Server

Objective:
	Write a simple Go web server that responds with a JSON object. This is a very common task for a backend developer.

Rules/Logic:

	1. Define a struct named ResponseData with a few fields, for example, Message (string) and Status (int).

	2. Create an HTTP handler function jsonHandler(w http.ResponseWriter, r *http.Request).

	3. Inside this handler:

		- Create an instance of your ResponseData struct (e.g., data := ResponseData{Message: "Success", Status: 200}).

		- Use json.Marshal to convert this struct into a JSON byte slice.

		- Important: Before writing the response, you must set the Content-Type header to application/json. This tells the client (like a browser) what kind of data it's receiving. You do this with: w.Header().Set("Content-Type", "application/json").

		- Write the JSON byte slice to the http.ResponseWriter.

	4. In your main function, register your jsonHandler for the path /api and start the server on port :8080.

---Terminal Output---
	// Test case 1: Accessing the root endpoint
	$ curl http://localhost:8080/
	Welcome to the home page!

	// Test case 2: Accessing the JSON API endpoint
	$ curl http://localhost:8080/api
	{"message":"Success","status":200}

Hints:
	- Remember to import the necessary packages: net/http, encoding/json, and fmt.
	- Always check for errors when marshalling JSON and handle them appropriately.
	- Use http.ListenAndServe(":8080", nil) to start your server.

This exercise will help you understand how to create a basic web server in Go that serves JSON responses, a fundamental skill for backend development.

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"` // It's good practice to add JSON tags
	Status  int    `json:"status"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!\n")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := ResponseData{Message: "Success", Status: 200}
	// IMPORTANT: Always check for errors.
	jsonData, err := json.Marshal(data)
	if err != nil {
		// If there's an error, send a server error response back to the client.
		http.Error(w, "Could not marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	// write a newline character to clean up the curl output
	w.Write([]byte("\n"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api", jsonHandler)
	fmt.Println("✅ Server starting...")
	fmt.Println("   Listening on port: 8080")
	fmt.Println("   Available endpoints:")
	fmt.Println("     - http://localhost:8080/")
	fmt.Println("     - http://localhost:8080/api")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
