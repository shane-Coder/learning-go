package main

import (
	"fmt"
	"strings"
)

func reverseWord(word string) string {
	list := []rune(word)
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return string(list)
}

func decipherMessage(message string) string {
	list := strings.Fields(message)
	var builder strings.Builder
	for i, word := range list {
		decodedWord := reverseWord(word)
		builder.WriteString(decodedWord)
		if i < len(list)-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func main() {
	encodedMessage1 := "olleH ,dlroW"
	encodedMessage2 := "oG si !gnizama"
	encodedMessage3 := "év café 🚀" // A string with multi-byte characters

	decodedMessage1 := decipherMessage(encodedMessage1)
	fmt.Println(decodedMessage1)

	decodedMessage2 := decipherMessage(encodedMessage2)
	fmt.Println(decodedMessage2)

	decodedMessage3 := decipherMessage(encodedMessage3)
	fmt.Println(decodedMessage3)
}
