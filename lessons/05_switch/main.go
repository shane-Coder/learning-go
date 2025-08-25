// package main

// import "fmt"

// func main() {
// 	day := "Friday"

// 	fmt.Printf("The day is %s. Here's the plan:\n", day)

// 	switch day {
// 	case "Monday", "Tuesday", "Wednesday", "Thursday":
// 		fmt.Println("It's a weekday. Time to learn Go!")
// 	case "Friday":
// 		fmt.Println("It's Friday! One last push before the weekend.")
// 	case "Saturday", "Sunday":
// 		fmt.Println("It's the weekend! Time to relax.")
// 	default:
// 		fmt.Println("That's not a valid day.")
// 	}
// }

// Advanced Switch statement code

package main

import "fmt"

func main() {
	score := 75

	fmt.Printf("The score is %d. Here is the remark:\n", score)

	// This is a "tagless switch". It's like an if/else if chain.
	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 80:
		fmt.Println("Very Good")
	case score >= 70:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Average")
	default:
		fmt.Println("Needs Improvement")
	}
}
