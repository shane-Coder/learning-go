package main

import "fmt"

func main() {
	// We initialize our counter outside the loop
	powerLevel := 3

	fmt.Println("Powering up...!")

	// The loop will run as long as powerLevel is greater than 0
	for powerLevel > 0 {
		fmt.Printf("Power level at: %d\n", powerLevel)
		// We decrement the counter inside the loop
		powerLevel--
	}

	fmt.Println("Liftoff!")
}
