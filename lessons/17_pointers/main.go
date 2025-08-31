package main

import "fmt"

func main() {
	// Create a normal variable holding a value.
	originalValue := 42
	fmt.Printf("1. Original variable 'originalValue': %d\n", originalValue)

	// Create a pointer that stores the memory address of 'originalValue'.
	// The type of a pointer to an int is *int.
	pointerToValue := &originalValue
	fmt.Printf("2. The pointer 'pointerToValue' holds the memory address: %v\n", pointerToValue)

	// Use the pointer to read the value. This is called dereferencing.
	fmt.Printf("3. Reading the value using the pointer: %d\n", *pointerToValue)

	// Use the pointer to CHANGE the value.
	fmt.Println("--- Changing the value through the pointer ---")
	*pointerToValue = 100

	// The original variable has now been changed!
	fmt.Printf("4. The original variable 'originalValue' is now: %d\n", originalValue)
}
