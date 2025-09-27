package main

import (
	"errors"
	"fmt"
)

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	value, err := factorial(n - 1)
	if err != nil {
		return 0, err
	}
	return n * value, nil
}

func main() {
	result, err := factorial(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial of 5 is:", result)
	}

	result, _ = factorial(0)
	fmt.Println("Factorial of 0 is:", result)

	_, err = factorial(-2)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
