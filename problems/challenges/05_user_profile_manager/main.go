package main

import "fmt"

type Profile struct {
	Username string
	Age      int
}

type ValidationError struct {
	Message string
}

type DuplicateUserError struct {
	Username string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", e.Message)
}

func (e *DuplicateUserError) Error() string {
	return fmt.Sprintf("Duplicate User Error: username '%s' already exists", e.Username)
}

// It's better practice to keep the database inside a struct, but a global variable is fine for this simple challenge.
var database = make(map[string]Profile)

func createProfile(username string, age int) (*Profile, error) {
	if len(username) < 3 {
		return nil, &ValidationError{Message: "username must be at least 3 characters long"}
	}
	if age < 18 {
		return nil, &ValidationError{Message: "age must be at least 18"}
	}

	if _, ok := database[username]; ok {
		return nil, &DuplicateUserError{Username: username}
	}

	profile := Profile{
		Username: username,
		Age:      age,
	}
	database[username] = profile

	return &profile, nil
}

func main() {
	_, err := createProfile("shivam", 25)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Profile for 'shivam' created successfully.")
	}

	_, err = createProfile("shivam", 30) // Test duplicate
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	_, err = createProfile("ab", 22) // Test validation
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	_, err = createProfile("alice", 17) // Test age validation
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
