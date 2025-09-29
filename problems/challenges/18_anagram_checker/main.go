package main

import (
	"fmt"
	"strings"
	"unicode"
)

// A helper function to create a clean, normalized string.
func normalize(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

func areAnagrams(s1, s2 string) bool {
	s1 = normalize(s1)

	s2 = normalize(s2)

	if len(s1) != len(s2) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range s1 {
		count[ch] += 1
	}
	for _, ch := range s2 {
		count[ch] -= 1
	}
	for _, ct := range count {
		if ct != 0 {
			return false // If any count is not zero, they are not anagrams.
		}
	}
	return true
}

func main() {
	fmt.Println(areAnagrams("listen", "silent"))
	fmt.Println(areAnagrams("hello", "world"))
	fmt.Println(areAnagrams("Dormitory", "dirty room##"))
	fmt.Println(areAnagrams("Go", "go go"))
}
