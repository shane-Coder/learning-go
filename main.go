package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInt := 100
	myFloat := 99.52873

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
