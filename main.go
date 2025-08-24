package main

import "fmt"

func main() {
	// Let's check if someone is old enough to vote.
	age := 17

	fmt.Printf("Checking voting eligibility for age: %v\n", age)

	// The 'if' statement checks the condition.
	if age >= 18 {
		// This block only runs if age is 18 or greater.
		fmt.Println("Result: You are eligible to vote!")
	} else {
		// This block only runs if age is less than 18.
		fmt.Println("Result: You are not eligible to vote yet.")
	}

	fmt.Println("---------------------------------")
	fmt.Println("The program has finished checking.")
}
