/*
Problem Statement #11: String Utilities Package

Objective:
	Create a reusable local package named stringutils with helper functions for string manipulation. Then, create a main program that imports and uses this package.

Rules/Logic:

	1. In the root of your learning-go project, create a new folder (package) named stringutils.

	2. Inside the stringutils folder, create a file named stringutils.go.

	3. In this new file, define two exported functions:

		1. Reverse(s string) string: This function should take a string and return a new string with the characters in reverse order.

		2. ToUpper(s string) string: This function should take a string and return its uppercase version. (Hint: you can use the ToUpper function from the standard library's strings package).

	4. In your main program, import your new local package (learning-go/stringutils).

	5. Call both Reverse and ToUpper with sample strings and print the results to the console.

---Terminal Output---
	// Test case 1
	result of test case1:
	dlroW olleH
	HELLO WORLD

	// Test case 2
	result of test case2:
	dlrow eht ot emoclew
	WELCOME TO THE WORLD
*/

package main

import (
	"fmt"
	"learning-go/stringutils" // importing the local helper package
)

func main() {
	// test case 1
	string1 := "Hello World"
	rev := stringutils.Reverse(string1)
	up := stringutils.ToUpper(string1)
	fmt.Printf("result of test case1:\n%s\n%s\n\n", rev, up)

	// test case 2
	string2 := "welcome to the world"
	rev = stringutils.Reverse(string2)
	up = stringutils.ToUpper(string2)
	fmt.Printf("result of test case2:\n%s\n%s\n\n", rev, up)
}
