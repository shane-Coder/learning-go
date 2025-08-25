// package main

// import "fmt"

// func main() {
// 	// Let's check if someone is old enough to vote.
// 	age := 20

// 	fmt.Printf("Checking voting eligibility for age: %v\n", age)

// 	// The 'if' statement checks the condition.
// 	if age >= 18 {
// 		// This block only runs if age is 18 or greater.
// 		fmt.Println("Result: You are eligible to vote!")
// 	} else {
// 		// This block only runs if age is less than 18.
// 		fmt.Println("Result: You are not eligible to vote yet.")
// 	}

// 	fmt.Println("---------------------------------")
// 	fmt.Println("The program has finished checking.")
// }

package main

import "fmt"

func main() {
	// --- Example 1: Using if, else if, and else ---
	score := 85

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
	age := 25
	day := "Monday"

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
