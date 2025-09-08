package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	now := time.Now()
	fmt.Println("Current time is:", now)

	// Format the time into a custom string using the reference layout
	// Layout: "Year-Month-Day Hour:Minute:Second"
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("Formatted time: %s\n", formattedTime)

	// You can use other standard layouts
	fmt.Printf("Standard Kitchen format: %s\n", now.Format(time.Kitchen))

	fmt.Println("\n--- Parsing a time string ---")
	// Define the layout the string is in
	layout := "2006-01-02"
	dateString := "2025-12-25"

	// Parse the string into a time.Time object
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Printf("Parsed time: %v\n", parsedTime)
		fmt.Printf("It's a %v!\n", parsedTime.Weekday())
	}
}
