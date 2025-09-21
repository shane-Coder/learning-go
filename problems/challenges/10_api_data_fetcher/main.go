package main

import (
	"encoding/json"
	"fmt"
	"io" // We need the io package for ReadAll
	"log"
	"net/http" // The package for making HTTP requests
)

// 1. The Todo struct with JSON tags.
// This is exactly like the structs from our JSON parser challenge.
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// fetchTodo is our main logic function.
func fetchTodo(todoID int) (*Todo, error) {
	// 2. Construct the URL for the API endpoint.
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", todoID)

	// 3. Make the HTTP GET request.
	// http.Get is the simplest way to do this.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 4. CRITICAL: Always close the response body when you're done with it.
	defer resp.Body.Close()

	// 5. Read the entire response body into a byte slice.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 6. Unmarshal the JSON data into our Todo struct.
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, err
	}

	// 7. Return the populated struct.
	return &todo, nil
}

func main() {
	// Call our function to fetch the todo with ID 1.
	todo, err := fetchTodo(1)
	if err != nil {
		log.Fatalf("Failed to fetch todo: %v", err)
	}

	// Print the formatted output as requested.
	fmt.Println("--- Fetched To-Do Item ---")
	fmt.Printf("ID: %d\n", todo.ID)
	fmt.Printf("User ID: %d\n", todo.UserID)
	fmt.Printf("Title: %s\n", todo.Title)
	fmt.Printf("Completed: %t\n", todo.Completed)
}
