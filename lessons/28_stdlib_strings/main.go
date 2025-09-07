package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- `strings` Package ---")
	myString := "hello, world! hello, go!"

	// Check if a string contains a substring
	fmt.Printf("Does '%s' contain 'world'? %t\n", myString, strings.Contains(myString, "world"))

	// Replace all occurrences of a substring
	replacedString := strings.ReplaceAll(myString, "hello", "hola")
	fmt.Printf("String after replacement: %s\n", replacedString)

	// Split a string into a slice
	words := strings.Split(myString, " ")
	fmt.Printf("String split into a slice: %#v\n", words)

	// Join a slice of strings into a single string
	joinedString := strings.Join([]string{"learning", "go", "is", "fun"}, " ")
	fmt.Printf("Slice joined into a string: %s\n", joinedString)

	fmt.Println("\n--- `strconv` Package ---")
	myNumberString := "123"
	// Convert string to integer (Atoi)
	myInt, err := strconv.Atoi(myNumberString)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("'%s' as an integer is: %d\n", myNumberString, myInt)
	}

	// Convert integer to string (Itoa)
	myInteger := 456
	myNewString := strconv.Itoa(myInteger)
	fmt.Printf("%d as a string is: '%s'\n", myInteger, myNewString)
}
