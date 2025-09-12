package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "!", "")
	list := strings.Fields(text)
	count := make(map[string]int)
	for _, word := range list {
		count[word] += 1
	}
	return count
}

func main() {
	text := "Hello world, this is a test. Hello world again!"
	result := countWords(text)
	fmt.Printf("Word frequencies: %v\n", result)
}
