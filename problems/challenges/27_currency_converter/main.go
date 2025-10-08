package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Converter struct {
	rates map[string]float64
}

func NewConverter() *Converter {
	var rates = map[string]float64{
		"USD": 1.0,
		"EUR": 0.93,
		"INR": 83.30,
	}
	return &Converter{rates: rates}
}

// func (c *Converter) Convert(query string) (float64, error) {
// 	list := strings.SplitN(query, " ", 3)
// 	if len(list) != 3 {
// 		return 0, fmt.Errorf("malformed query")
// 	}
// 	if _, ok := c.rates[list[1]]; !ok {
// 		return 0, fmt.Errorf("unknown currency to have: %s", list[1])
// 	}
// 	if _, ok := c.rates[list[2]]; !ok {
// 		return 0, fmt.Errorf("unknown currency to convert to: %s", list[2])
// 	}
// 	amount, err := strconv.ParseFloat(list[0], 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("invalid amount: %s", list[0])
// 	}
// 	rate_of_A := c.rates[list[1]]
// 	rate_of_B := c.rates[list[2]]
// 	result := (amount / rate_of_A) * rate_of_B
// 	return result, nil
// }

func (c *Converter) Convert(query string) (float64, error) {
	parts := strings.SplitN(query, " ", 3)
	if len(parts) != 3 {
		// More descriptive error message from the challenge description
		return 0, fmt.Errorf("invalid query format. expected 'AMOUNT FROM_CURRENCY TO_CURRENCY'")
	}

	amount, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %s", parts[0])
	}

	// --- START OF IMPROVEMENT ---
	fromCurrency := strings.ToUpper(parts[1])
	toCurrency := strings.ToUpper(parts[2])
	// --- END OF IMPROVEMENT ---

	rateA, okA := c.rates[fromCurrency]
	if !okA {
		return 0, fmt.Errorf("currency '%s' not found", fromCurrency)
	}

	rateB, okB := c.rates[toCurrency]
	if !okB {
		return 0, fmt.Errorf("currency '%s' not found", toCurrency)
	}

	result := (amount / rateA) * rateB
	return result, nil
}

func main() {
	converter := NewConverter()

	// Test case 1: USD to INR
	result1, err1 := converter.Convert("100 USD INR")
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("100 USD is %.2f INR\n", result1)
	}

	// Test case 2: EUR to USD
	result2, err2 := converter.Convert("50 EUR USD")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Printf("50 EUR is %.2f USD\n", result2)
	}

	// Test case 3: Unknown currency
	_, err3 := converter.Convert("100 JPY INR")
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

	// Test case 4: Malformed query
	_, err4 := converter.Convert("100 USD")
	if err4 != nil {
		fmt.Println("Error:", err4)
	}

	// Test case 5: Invalid amount
	_, err5 := converter.Convert("abc USD INR")
	if err5 != nil {
		fmt.Println("Error:", err5)
	}

	// Test case 6: Valid conversion
	_, err6 := converter.Convert("200 EUR YEN")
	if err6 != nil {
		fmt.Println("Error:", err6)
	}
}
