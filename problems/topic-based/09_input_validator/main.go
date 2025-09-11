/*
Problem Statement #9: User Input Validator

Objective:
	Write a Go program with a function that validates a username string and returns an error if the validation fails.

Rules/Logic:

	1. Create a function validateUsername(username string) error.

	2. This function will return an error if the username is invalid and nil if it is valid.

	3. The validation rules are:

		1. If the username is an empty string (""), return an error with the message "username cannot be empty".

		2. If the length of the username is greater than 10 characters, return an error with the message "username is too long".

	4. In your main function, call validateUsername with several test cases (a valid name, an empty name, and a name that is too long).

	5. Use the standard if err != nil pattern to check the returned error and print either a success message or the specific error message.

---Terminal Output---
	// Test Case 1: A valid name ---
		input is Hanuman
		output is Name is a valid name
	// Test Case 2: An empty name ---
		input is ""
		output is An error occurred: username cannot be empty
	// Test Case 3: A name that is too long---
		input is Kumbhakarna
		output is An error occurred: username is too long
*/

package main

import (
	"errors"
	"fmt"
)

func validateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	if len(username) > 10 {
		return errors.New("username is too long")
	}
	return nil
}

func main() {
	fmt.Println("--- Test Case 1: A valid name ---")
	err := validateUsername("Hanuman")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Name is a valid name\n")
	}

	fmt.Println("--- Test Case 2: An empty name ---")
	err = validateUsername("")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Name is a valid name\n")
	}

	fmt.Println("--- Test Case 3: A name that is too long---")
	err = validateUsername("Kumbhakarna")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Name is a valid name\n")
	}
}
