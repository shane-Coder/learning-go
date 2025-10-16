package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	store := make(map[rune]int)
	start := 0
	size := 0
	for end, char := range s {

		// logic 1
		// if store[char] == 1 {
		// 	start = end
		// } else {
		// 	store[char] += 1
		// }
		// if end-start+1 > size {
		// 	size = end - start + 1
		// }

		// logic 2
		// Check if we have seen this character before.
		if lastSeenIndex, ok := store[char]; ok {
			// If we have, we need to check if the last time we saw it
			// is within our current window (i.e., its index is >= start).
			if lastSeenIndex >= start {
				// If it is, we have a repeat inside our window.
				// We must move the start of our window to the position
				// right AFTER the last time we saw this character.
				start = lastSeenIndex + 1
			}
		}
		// Update the last seen index for the current character to its current position.
		store[char] = end

		// Calculate the length of the current valid window.
		currentLength := end - start + 1

		// Update our maximum length if the current one is bigger.
		if currentLength > size {
			size = currentLength
		}
	}
	return size
}

func main() {
	s1 := "abcabcbb"
	fmt.Printf("Length of longest substring in '%s': %d\n", s1, lengthOfLongestSubstring(s1))

	s2 := "bbbbb"
	fmt.Printf("Length of longest substring in '%s': %d\n", s2, lengthOfLongestSubstring(s2))

	s3 := "pwwkew"
	fmt.Printf("Length of longest substring in '%s': %d\n", s3, lengthOfLongestSubstring(s3))

	s4 := "pwwkew"
	fmt.Printf("Length of longest substring in '%s': %d\n", s4, lengthOfLongestSubstring(s4))

	s5 := "abcdeafg"
	fmt.Printf("Length of longest substring in '%s': %d\n", s5, lengthOfLongestSubstring(s5))

	s6 := "abba"
	fmt.Printf("Length of longest substring in '%s': %d\n", s6, lengthOfLongestSubstring(s6))
}
