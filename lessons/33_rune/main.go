// learning about runes in Go

// package main

// import "fmt"

// func main() {
// 	myString := "नमस्ते"

// 	fmt.Println("--- Looping with Bytes (Incorrect) ---")
// 	// This loop gives us the raw bytes. Notice how the length is 18!
// 	fmt.Printf("Length of string in bytes: %d\n", len(myString))
// 	for i := 0; i < len(myString); i++ {
// 		fmt.Printf("Byte %d: %c\n", i, myString[i]) // This will print garbage for some bytes
// 	}

// 	fmt.Println("\n--- Looping with Runes (Correct) ---")
// 	// A for...range loop on a string gives us runes. Notice the length is 6!
// 	runeSlice := []rune(myString)
// 	fmt.Printf("Length of string in runes (characters): %d\n", len(runeSlice))
// 	for index, r := range myString {
// 		fmt.Printf("Rune at byte position %d: %c\n", index, r)
// 	}
// }

// practice with a problem
// To practice this, create a function countCharacters(text string) int that takes a string (which might contain multi-byte characters like emojis or other languages) and returns the correct number of characters in it.
// This will help solidify your understanding of len() vs. using runes.

package main

import (
	"fmt"
	"unicode/utf8"
)

func countCharacters(text string) int {
	list := []rune(text)
	ct := 0
	for index, ch := range list {
		// size of each character in bytes
		fmt.Printf("this character index, value and size: %d %c %d\n", index, ch, utf8.RuneLen(ch))
		ct += 1
	}
	return ct
	// return len([]rune(text))
	// return utf8.RuneCountInString(text)
}

func main() {
	testString := "Hello, दुनिया! 🚀"
	fmt.Printf("Input: %s\n", testString)
	fmt.Printf("Byte count (len): %d\n", len(testString))
	fmt.Printf("Character count (rune): %d\n", countCharacters(testString))
}
