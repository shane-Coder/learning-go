package main

import (
	"errors"
	"fmt"
)

// A function that returns a result AND an error.
func divide(a, b float64) (float64, error) {
	// Check for an invalid input that would cause an error.
	if b == 0.0 {
		// If it fails, return a zero value for the result and a new error.
		return 0.0, errors.New("cannot divide by zero")
	}
	// If it succeeds, return the result and a 'nil' error.
	return a / b, nil
}

func main() {
	fmt.Println("--- Test Case 1: Successful Division ---")
	result, err := divide(0.0, 0.0)

	// This is the standard Go error check.
	if err != nil {
		// This block is skipped because 'err' is nil.
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	fmt.Println("\n--- Test Case 2: Failed Division ---")
	result, err = divide(10.0, 0.1)

	if err != nil {
		// This block runs because the returned 'err' is not nil.
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
