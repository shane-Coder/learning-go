package main

import "fmt"

func main() {
	day := "Friday"

	fmt.Printf("The day is %s. Here's the plan:\n", day)

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a weekday. Time to learn Go!")
	case "Friday":
		fmt.Println("It's Friday! One last push before the weekend.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend! Time to relax.")
	default:
		fmt.Println("That's not a valid day.")
	}
}
