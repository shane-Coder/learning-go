package main

import "fmt"

func main() {
	// 1. Using the short declaration operator `:=`
	name := "Shivam"
	age := 25
	isLearning := true
	score := 95.5

	// 2. Using the `var` keyword
	var location string = "Banda, Uttar Pradesh"

	// Let's print these values to the console
	fmt.Println("--- My Information ---")

	// fmt.Printf allows for formatted printing.
	// %v is a placeholder for the value.
	// %T is a placeholder for the type.
	// \n creates a new line.
	fmt.Printf("Name: %v, Type: %T\n", name, name)
	fmt.Printf("Age: %v, Type: %T\n", age, age)
	fmt.Printf("Location: %v, Type: %T\n", location, location)
	fmt.Printf("Currently Learning Go? %v, Type: %T\n", isLearning, isLearning)
	fmt.Printf("Score: %v, Type: %T\n", score, score)
}
