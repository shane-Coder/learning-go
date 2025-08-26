package main

import "fmt"

func main() {
	counter := 0

	// This is an infinite loop
	for {
		// We increment the counter at the start of each loop
		counter++

		// If the counter is an odd number, skip this iteration
		if counter%2 != 0 {
			continue // Jump to the next loop run
		}

		fmt.Printf("Counter is at an even number: %d\n", counter)

		// If the counter reaches 10, exit the loop
		if counter >= 10 {
			break // Exit the loop immediately
		}
	}

	fmt.Println("Loop has been broken. Program finished.")
}
