package main

import "fmt"

func main() {
	// Creating a slice of strings
	fruits := []string{"Apple", "Banana", "Cherry", "Mango", "Pineapple"}
	fmt.Println("My fruits:", fruits)

	// Getting the length of the slice
	fmt.Printf("I have %d fruits.\n", len(fruits))

	// Accessing a specific element (index starts at 0)
	fmt.Println("The first fruit is:", fruits[3])

	fmt.Println("---------------------------------")

	// Using a for...range loop to iterate over the slice
	fmt.Println("Here is a list of my fruits:")
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}
