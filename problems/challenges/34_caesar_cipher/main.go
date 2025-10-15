package main

import (
	"fmt"
	"strings"
)

// logic 1
// func caesarCipher(text string, shift int) string {
// 	result := ""
// 	for _, char := range text {
// 		if char >= 'a' && char <= 'z' {
// 			result += string(rune((int(char)-int('a')+shift)%26 + int('a')))
// 		} else if char >= 'A' && char <= 'Z' {
// 			result += string(rune((int(char)-int('A')+shift)%26 + int('A')))
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	return result
// }

// logic 2
func caesarCipher(text string, shift int) string {
	// Use a strings.Builder for efficient string construction.
	var builder strings.Builder
	shiftRune := rune(shift) // Convert shift to a rune for calculations

	for _, char := range text {
		// Check for lowercase letters
		if char >= 'a' && char <= 'z' {
			newChar := 'a' + (char-'a'+shiftRune)%26
			builder.WriteRune(newChar)
		} else if char >= 'A' && char <= 'Z' { // Check for uppercase letters
			newChar := 'A' + (char-'A'+shiftRune)%26
			builder.WriteRune(newChar)
		} else {
			// If it's not a letter, add it unchanged.
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func main() {
	text1 := "Hello, World!"
	shift1 := 3
	encrypted1 := caesarCipher(text1, shift1)
	fmt.Printf("Original: '%s'\nEncrypted: '%s'\n\n", text1, encrypted1)

	text2 := "Zebra-123"
	shift2 := 5
	encrypted2 := caesarCipher(text2, shift2)
	fmt.Printf("Original: '%s'\nEncrypted: '%s'\n", text2, encrypted2)
}
