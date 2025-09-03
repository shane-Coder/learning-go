// This file is part of the 'main' package.
package main

import (
	"fmt"
	"learning-go/calculator" // Import our local package
)

func main() {
	sum := calculator.Add(10, 5)
	fmt.Printf("The sum is: %d\n", sum)

	// This next line would cause an error because 'subtract' is not exported:
	// diff := calculator.subtract(10, 5)
}
