package main

import "fmt"

func main() {
	// --- Example 1: Using if, else if, and else ---
	score := 10

	fmt.Printf("The score is %d. Calculating grade...\n", score)

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	fmt.Println("---------------------------------")

	// --- Example 2: Using Logical Operators (&& and ||) ---
	age := 10
	day := "Sunday"

	fmt.Printf("Checking work status for age %d on a %s...\n", age, day)

	// Using AND (&&)
	if age >= 18 && age <= 65 {
		fmt.Println("You are of standard working age.")
	}

	// Using OR (||)
	if day == "Saturday" || day == "Sunday" {
		fmt.Println("It's the weekend! Time to rest.")
	} else {
		fmt.Println("It's a weekday. Back to work or learning Go!")
	}
}
