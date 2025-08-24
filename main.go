package main

import "fmt"

func main() {
	// 1. Declaring a Constant
	const Pi = 3.14159

	fmt.Printf("The value of Pi is always: %v\n", Pi)
	fmt.Println("---------------------------------")

	// 2. Arithmetic Operations with Integers
	a := 20
	b := 10

	fmt.Println("--- Integer Operations ---")
	fmt.Printf("%v + %v = %v\n", a, b, a+b) // Addition
	fmt.Printf("%v - %v = %v\n", a, b, a-b) // Subtraction
	fmt.Printf("%v * %v = %v\n", a, b, a*b) // Multiplication
	fmt.Printf("%v / %v = %v\n", a, b, a/b) // Division
	fmt.Println("---------------------------------")

	// 3. String Concatenation
	firstName := "Shivam"
	lastName := "Omer"

	// We can join strings with the '+' operator
	fullName := firstName + " " + lastName

	fmt.Println("--- String Operations ---")
	fmt.Printf("Full Name: %v\n", fullName)
}
