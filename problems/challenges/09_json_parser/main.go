package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define the structs to match the JSON structure.
// type SocialProfile struct {
// 	Platform string
// 	Username string
// }

// type User struct {
// 	ID       int
// 	Name     string
// 	Email    string
// 	Profiles []SocialProfile
// }

// Structs with JSON tags for proper unmarshalling
type SocialProfile struct {
	Platform string `json:"platform"`
	Username string `json:"username"`
}

type User struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Profiles []SocialProfile `json:"profiles"`
}

func parseUserJSON(jsonString string) (User, error) {
	jsonInput := []byte(jsonString)
	var user User
	err := json.Unmarshal(jsonInput, &user)
	return user, err
}

func main() {
	const userJSON = `{
	    "id": 101,
	    "name": "Shivam",
	    "email": "shivam@example.com",
	    "profiles": [
		{
		    "platform": "GitHub",
		    "username": "shane-Coder"
		},
		{
		    "platform": "LinkedIn",
		    "username": "programmer-shivam"
		}
	    ]
	}`
	user, err := parseUserJSON(userJSON)
	if err != nil {
		// It's good practice to log the actual error
		log.Fatalf("Error parsing JSON: %v", err)
	}
	// simple confirmation message
	// fmt.Printf("successful created profile: %v\n", user)
	// Detailed output
	fmt.Println("--- User Profile ---")
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Println("Social Profiles:")
	for _, profile := range user.Profiles {
		fmt.Printf("  - Platform: %s, Username: %s\n", profile.Platform, profile.Username)
	}

}
