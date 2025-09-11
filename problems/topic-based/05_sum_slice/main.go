/*
Problem Statement #5: Sum a Slice

Objective:
	Write a Go program with a dedicated function that calculates the sum of all numbers in a slice of integers.

Rules/Logic:

	1. Your main function should define a slice of integers (e.g., numbers := []int{10, 20, 30, 5}).

	2. You must create a separate function named sumSlice.

	3. The sumSlice function must accept one argument: a slice of integers ([]int).

	4. The sumSlice function must return one value: the calculated sum as an int.

	5. Inside the sumSlice function, use a for...range loop to add up all the numbers in the slice.

	6. Call your sumSlice function from main and print the returned result to the console.

---Terminal Output---
	// Test Case
	// Input: []int{10, 20, 30, 5}
	// Expected Output: 65
*/

package main

import "fmt"

func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	fmt.Println("Code logic here")
	numbers := []int{10, 20, 30, 5}
	total := sumSlice(numbers)
	fmt.Printf("%d\n", total)
}
