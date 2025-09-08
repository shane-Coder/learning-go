package main

import (
	"encoding/json"
	"fmt"
)

// We use struct field tags to control the JSON output.
type User struct {
	// The JSON key will be "fullName" instead of "Name".
	Name string `json:"fullName"`
	// The JSON key will be "age".
	Age int `json:"age"`
	// If this field is empty, it will be omitted from the JSON.
	Email string `json:"email,omitempty"`
	// This field will always be ignored by the JSON package.
	Password string `json:"-"`
}

func main() {
	// --- Marshalling (Go struct to JSON) ---
	fmt.Println("--- Marshalling ---")
	user1 := User{
		Name:     "Shivam Omer",
		Age:      25,
		Password: "a-secret-password",
	}

	// json.Marshal returns a byte slice and an error.
	jsonData, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// We convert the byte slice to a string to print it.
	fmt.Println("JSON output:", string(jsonData))

	// --- Unmarshalling (JSON to Go struct) ---
	fmt.Println("\n--- Unmarshalling ---")
	jsonInput := []byte(`{"fullName":"Alice","age":30,"email":"alice@example.com"}`)

	var user2 User
	// json.Unmarshal takes the JSON byte slice and a pointer to the struct
	// where the data should be stored.
	err = json.Unmarshal(jsonInput, &user2)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("Unmarshalled User: %+v\n", user2)
}
