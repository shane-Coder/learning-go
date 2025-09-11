/*
Problem Statement #6: Name Frequency Counter

Objective:
	Write a Go program with a dedicated function that takes a slice of names (strings) and returns a map counting how many times each name appears.

Rules/Logic:

	1. Your main function should define a slice of names (e.g., names := []string{"Shivam", "Alice", "Bob", "Shivam", "Alice", "Shivam"}).

	2. Create a separate function called countNames.

	3. The countNames function must accept one argument: a slice of strings ([]string).

	4. The countNames function must return a map[string]int, where the keys are the names and the values are their counts.

	5. Inside countNames, create an empty map. Loop over the input slice. For each name, increment its count in your map.

	6. Call your countNames function from main and print the resulting map.

---Terminal Output---
	// Test Case
	// Input: []string{"Shivam", "Alice", "Bob", "Shivam", "Alice", "Shivam"}
	// Expected Output: map[Alice:2 Bob:1 Shivam:3]
*/

package main

import "fmt"

func countNames(names []string) map[string]int {
	count := make(map[string]int) // method 1 to initialize map
	// count := map[string]int{"": 0} //method 2 to initialize map
	for _, name := range names {
		count[name] += 1
	}
	return count
}

func main() {
	fmt.Println("code logic here")
	names := []string{"Shivam", "Alice", "Bob", "Shivam", "Alice", "Shivam"}
	fmt.Printf("%v \n", countNames(names)) // method1 to solve the problem
	// count := countNames(names) // method2 to solve the problem
	// for name, freq := range count {
	// 	fmt.Printf("%s: %d\n", name, freq)
	// }
}
