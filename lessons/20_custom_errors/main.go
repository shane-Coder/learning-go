package main

import (
	"fmt"
)

// 1. Define a custom error struct.
type ValidationError struct {
	Code    int
	Message string
}

// 2. Implement the Error() method for the struct.
// This makes our ValidationError satisfy the built-in 'error' interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error (%d): %s", e.Code, e.Message)
}

// 3. Create a function that returns our custom error.
func validateAge(age int) error {
	if age < 0 {
		// Return an instance of our custom error.
		return &ValidationError{Code: 400, Message: "age cannot be negative"}
	}
	if age < 18 {
		return &ValidationError{Code: 403, Message: "user is underage"}
	}
	return nil
}

func main() {
	// --- Test Case 1: Underage Error ---
	err := validateAge(15)
	if err != nil {
		fmt.Println(err)
	}

	// --- Test Case 2: Negative Age Error ---
	err = validateAge(-5)
	if err != nil {
		fmt.Println(err)
	}

	// --- Test Case 3: No Error ---
	err = validateAge(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Age is valid.")
	}
}
