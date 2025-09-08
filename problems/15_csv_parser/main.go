/*
Problem Statement #15: CSV Data Parser

Objective:
	Write a Go program with a function that takes a single line of CSV (Comma Separated Values) data as a string, parses it, and returns a structured User object.

Rules/Logic:

	1. Define a struct named User with the following fields: Name (string), Age (int), and IsActive (bool).

	2. Create a function parseCSVLine(line string) (User, error) that takes a string like "shivam,25,true" as input.

	3. Inside the function:

		- Use strings.Split(line, ",") to split the string into a slice of parts.

		- Check if the slice has exactly 3 parts. If not, return an error.

		- Use strconv.Atoi() to convert the age part from a string to an integer.

		- Use strconv.ParseBool() to convert the active status part from a string to a boolean.

		- Handle any potential errors from these conversion functions.

		- Create a User struct with the parsed data and return it.

	4. In your main function, call parseCSVLine with a few test cases (a valid line and some invalid lines) and print either the resulting User struct or the error message.

---Terminal Output---
	// Testcase 1
	Input: "shivam, 25, true" -> Success: {Name:shivam Age:25 IsActive:true}
	// Testcase 2
	Input: "shivam,  26  ,  false  " -> Success: {Name:shivam Age:26 IsActive:false}
	// Testcase 3
	Input: "24,false" -> Error: invalid number of fields, expected 3
	// Testcase 4
	Input: "alice,thirty,true" -> Error: age field is invalid: strconv.Atoi: parsing "thirty": invalid syntax
	// Testcase 5
	Input: "bob,30,yes" -> Error: isActive field is invalid: strconv.ParseBool: parsing "yes": invalid syntax
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func parseCSVLine(line string) (User, error) {
	csvList := strings.Split(line, ",")
	if len(csvList) != 3 {
		return User{}, errors.New("invalid number of fields, expected 3")
	}
	// Trim whitespace before converting for robustness
	name := strings.TrimSpace(csvList[0])
	ageStr := strings.TrimSpace(csvList[1])
	activeStr := strings.TrimSpace(csvList[2])
	age, err1 := strconv.Atoi(ageStr)
	if err1 != nil {
		return User{}, fmt.Errorf("age field is invalid: %w", err1)
	}
	status, err2 := strconv.ParseBool(activeStr)
	if err2 != nil {
		return User{}, fmt.Errorf("isActive field is invalid: %w", err2)
	}
	return User{
		Name:     name,
		Age:      age,
		IsActive: status,
	}, nil
}

func main() {
	// Using a simple test helper function to keep main clean
	testLine := func(line string) {
		user, err := parseCSVLine(line)
		if err != nil {
			fmt.Printf("Input: \"%s\" -> Error: %v\n", line, err)
		} else {
			fmt.Printf("Input: \"%s\" -> Success: %+v\n", line, user)
		}
	}

	testLine("shivam, 25, true")
	testLine("shivam,  26  ,  false  ") // Test with whitespace
	testLine("24,false")                // Test invalid field count
	testLine("alice,thirty,true")       // Test invalid age
	testLine("bob,30,yes")              // Test invalid boolean
}
