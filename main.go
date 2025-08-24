package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --- Numeric Conversion ---
	myInteger := 42
	myFloat := 3.14

	// This would cause an error: invalid operation: myInteger + myFloat (mismatched types int and float64)
	// fmt.Println(myInteger + myFloat)

	// The CORRECT way: Convert the integer to a float64
	sum := float64(myInteger) + myFloat

	fmt.Println("--- Numeric Conversion ---")
	fmt.Printf("The sum is: %v, Type: %T\n", sum, sum)
	fmt.Println("---------------------------------")

	// --- String Conversion ---
	age := 25

	// We use strconv.Itoa() to convert the integer 'age' to a string
	ageAsString := strconv.Itoa(age)

	welcomeMessage := "My age is: " + ageAsString

	fmt.Println("--- String Conversion ---")
	fmt.Printf("The welcome message is: \"%v\"\n", welcomeMessage)
	fmt.Printf("Type of 'age' is %T, but type of 'ageAsString' is %T\n", age, ageAsString)
}
