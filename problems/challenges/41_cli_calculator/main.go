package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch {
	case operator == "+":
		return num1 + num2, nil
	case operator == "-":
		return num1 - num2, nil
	case operator == "*":
		return num1 * num2, nil
	case operator == "/":
		if num2 == 0 {
			// This is the specific error we need to handle.
			return 0, errors.New("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		// This handles any other operator, like 'x'.
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number1> <operator> <number2>")
		os.Exit(1)
	}
	number1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	number2, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: invalid number provided.")
		os.Exit(1)
	}
	result, err := calculate(number1, operator, number2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("%.2f %s %.2f = %.2f\n", number1, operator, number2, result)
}
