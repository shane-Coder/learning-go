package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice:", numbers)
	fmt.Println("---------------------------------")

	// Get a slice of the numbers from index 1 up to (but not including) index 4
	// This will be [1, 2, 3]
	subset := numbers[1:4]
	fmt.Println("Subset [1:4]:", subset)

	// Get a slice from the beginning up to index 5
	// This will be [0, 1, 2, 3, 4]
	firstFive := numbers[:5]
	fmt.Println("First Five [:5]:", firstFive)

	// Get a slice from index 5 to the end
	// This will be [5, 6, 7, 8, 9, 10]
	fromFiveUp := numbers[5:]
	fmt.Println("From Five Up [5:]:", fromFiveUp)
}
