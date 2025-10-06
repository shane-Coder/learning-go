package main

import (
	"fmt"
	"unicode"
)

func countCharacters(text string) map[rune]int {
	var result = make(map[rune]int)
	for _, ch := range text {
		if unicode.IsSpace(ch) {
			continue
		}
		result[ch] += 1
	}
	return result
}

func main() {
	text := "hello world 👋"
	frequencyMap := countCharacters(text)

	fmt.Println("--- Character Frequencies ---")
	for char, count := range frequencyMap {
		fmt.Printf("'%c': %d\n", char, count)
	}
}
