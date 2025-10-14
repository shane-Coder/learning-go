package main

import (
	"errors"
	"fmt"
)

func firstNonRepeating(s string) (rune, error) {
	result := make(map[rune]int)
	for _, char := range s {
		result[char] += 1
	}
	for _, char := range s {
		if result[char] == 1 {
			return char, nil
		}
	}
	return '0', errors.New("no non-repeating character found")
}

func main() {
	s1 := "swiss"
	char1, err1 := firstNonRepeating(s1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("First non-repeating char in '%s' is '%c'\n", s1, char1)
	}

	s2 := "aabbcc"
	_, err2 := firstNonRepeating(s2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	s3 := "hello"
	char3, err3 := firstNonRepeating(s3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Printf("First non-repeating char in '%s' is '%c'\n", s3, char3)
	}

	s4 := ""
	_, err4 := firstNonRepeating(s4)
	if err4 != nil {
		fmt.Println("Error:", err4)
	}
}
