package main

import (
	"fmt"
)

func romanToInt(s string) int {
	// 1. Create a map for Roman numeral values.
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	// 2. Loop through the string from left to right.
	for i := range n {
		currentValue := romanValues[rune(s[i])]

		// 3. Look ahead: Check if there is a next character and if its value is greater.
		if i+1 < n && romanValues[rune(s[i+1])] > currentValue {
			// Subtraction case (e.g., IV, IX, XL, XC, CD, CM)
			total -= currentValue
		} else {
			// Addition case (or it's the last character)
			total += currentValue
		}
	}

	return total
}

func main() {
	fmt.Println("III =", romanToInt("III"))
	fmt.Println("LVIII =", romanToInt("LVIII"))
	fmt.Println("MCMXCIV =", romanToInt("MCMXCIV"))
}
