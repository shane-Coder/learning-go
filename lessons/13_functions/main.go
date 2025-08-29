package main

import (
	"errors" // We need this package to create a new error
	"fmt"
)

// add takes two integers and returns their sum as an integer.
func add(a int, b int) int {
	return a + b
}

// divide takes two floats and returns a float AND an error.
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		// It's idiomatic to return a zero value for the result and an error.
		return 0, errors.New("cannot divide by zero")
	}
	// If there's no error, we return the result and 'nil' for the error.
	return a / b, nil
}

func main() {
	// --- Calling a simple function ---
	sum := add(5, 10)
	fmt.Printf("The sum is: %d\n", sum)
	fmt.Println("---------------------------------")

	// --- Calling a function that can return an error ---
	result, err := divide(10.20, 4.0)
	if err != nil {
		// An error occurred
		fmt.Printf("Error: %v\n", err)
	} else {
		// No error, everything is ok
		fmt.Printf("The result of division is: %.2f\n", result)
	}

	// --- Calling the function with an invalid input ---
	result, err = divide(10.0, 0.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("The result of division is: %.2f\n", result)
	}
}
