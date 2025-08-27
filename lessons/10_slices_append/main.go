package main

import "fmt"

func main() {
	// Start with an initial slice of fruits
	fruits := []string{"Apple", "Banana"}
	fmt.Println("Original fruits:", fruits)
	fmt.Printf("Original length: %d\n", len(fruits))
	fmt.Println("---------------------------------")

	// Add one new fruit to the slice
	// IMPORTANT: We must assign the result of append() back to the 'fruits' variable
	fruits = append(fruits, "Cherry")

	fmt.Println("After adding Cherry:", fruits)
	fmt.Printf("New length: %d\n", len(fruits))
	fmt.Println("---------------------------------")

	// Add multiple fruits at once
	fruits = append(fruits, "Mango", "Pineapple")

	fmt.Println("After adding more fruits:", fruits)
	fmt.Printf("Final length: %d\n", len(fruits))
}
