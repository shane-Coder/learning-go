/*
Problem Statement #4: Even Number Filter

Objective:
	Write a Go program that takes a slice of integers and creates a new slice containing only the even numbers from the original list.

Rules/Logic:

	1. Start with a pre-defined slice of numbers, for example: numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.

	2. Create a new, empty slice to store your results, for example: evens := []int{}.

	3. Use a for...range loop to iterate over the numbers slice.

	4. Inside the loop, use an if statement to check if the current number is even. (Hint: The modulo operator % is perfect for this. A number is even if number % 2 == 0).

	5. If a number is even, use the append() function to add it to your evens slice.

	6. After the loop finishes, print the final evens slice.

--- Terminal Output ---
	// Test case : Filtering even and odd numbers from a slice
	slice of numbers : [1 2 3 4 5 6 7 8 9 10]
	Even numbers: [2 4 6 8 10]
	Odd numbers: [1 3 5 7 9]
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := []int{}
	odd := []int{}
	for _, num := range numbers {
		if num&1 == 1 {
			odd = append(odd, num)
		} else {
			even = append(even, num)
		}
	}
	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)
}
