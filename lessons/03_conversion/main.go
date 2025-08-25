// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// --- Numeric Conversion ---
// 	myInteger := 42
// 	myFloat := 3.14

// 	// This would cause an error: invalid operation: myInteger + myFloat (mismatched types int and float64)
// 	// fmt.Println(myInteger + myFloat)

// 	// The CORRECT way: Convert the integer to a float64
// 	sum := float64(myInteger) + myFloat

// 	fmt.Println("--- Numeric Conversion ---")
// 	fmt.Printf("The sum is: %v, Type: %T\n", sum, sum)
// 	fmt.Println("---------------------------------")

// 	// --- String Conversion ---
// 	age := 25

// 	// We use strconv.Itoa() to convert the integer 'age' to a string
// 	ageAsString := strconv.Itoa(age)

// 	welcomeMessage := "My age is: " + ageAsString

// 	fmt.Println("--- String Conversion ---")
// 	fmt.Printf("The welcome message is: \"%v\"\n", welcomeMessage)
// 	fmt.Printf("Type of 'age' is %T, but type of 'ageAsString' is %T\n", age, ageAsString)
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInt := 100
	myFloat := 99.5

	// --- CORRECT way for Integers ---
	intAsString := strconv.Itoa(myInt)
	fmt.Printf("Integer %v is a %T, converted to \"%v\" which is a %T\n", myInt, myInt, intAsString, intAsString)

	// --- INCORRECT way for Floats ---
	// This line will cause a compile error: cannot use myFloat (variable of type float64) as int value in argument to strconv.Itoa
	// floatAsString_WRONG := strconv.Itoa(myFloat)

	// --- CORRECT way for Floats ---
	// We convert 99.5 to a string, keeping 2 decimal places of precision.
	floatAsString_RIGHT := strconv.FormatFloat(myFloat, 'f', 2, 64)
	fmt.Printf("Float %v is a %T, converted to \"%v\" which is a %T\n", myFloat, myFloat, floatAsString_RIGHT, floatAsString_RIGHT)
}
