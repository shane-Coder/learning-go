// examples/04_generics/main.go
package main

import (
	"errors"
	"fmt"
)

// --- Example 1: A Generic Function with 'any' ---
// This function is generic over any type 'T'.
// It's robust because it returns a second value, an error, to handle the empty slice case.

func getLastElement[T any](s []T) (T, error) {
	if len(s) == 0 {
		var zero T // Creates the "zero value" for type T (0 for int, "" for string, etc.)
		return zero, errors.New("cannot get last element of an empty slice")
	}
	return s[len(s)-1], nil
}

// --- Example 2: A Generic Function with Type Constraints ---

// Number is a type constraint (an interface) that permits any type
// whose underlying type is either int or float64.
type Number interface {
	~int | ~float64
}

// This function is generic over any type 'T' that satisfies the Number constraint.
// This allows us to safely use the '+' operator inside the function.
func sumSlice[T Number](s []T) T {
	var total T
	for _, value := range s {
		total += value
	}
	return total
}

func main() {
	fmt.Println("--- Generic Function with 'any' ---")
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"hello", "world"}
	emptySlice := []float64{}

	lastInt, _ := getLastElement(intSlice)
	fmt.Println("Last element of int slice:", lastInt)

	lastString, _ := getLastElement(stringSlice)
	fmt.Println("Last element of string slice:", lastString)

	_, err := getLastElement(emptySlice)
	fmt.Println("Error from empty slice:", err)

	fmt.Println("\n--- Generic Function with Constraints ---")
	floatSlice := []float64{1.1, 2.2, 3.3}
	fmt.Println("Sum of int slice:", sumSlice(intSlice))
	fmt.Printf("Sum of float slice: %.2f\n", sumSlice(floatSlice))
}
