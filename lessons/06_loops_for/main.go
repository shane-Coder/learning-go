package main

import "fmt"

func main() {
	fmt.Println("Loop is starting...")

	// init:      i := 0
	// condition: i < 5
	// post:      i++

	// 1st type for counting
	for i := 10; i > 0; i-- {
		fmt.Printf("This is iteration number: %d\n", i)
	}

	// 2nd type simple number of times loop
	for i := range 5 {
		fmt.Printf("This is iteration number: %d\n", i)
	}

	// For each item in the 'fruits' slice...
	// example for iteration
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("At index %d, we have a %s\n", index, fruit)
	}

	fmt.Println("Loop has finished.")
}
