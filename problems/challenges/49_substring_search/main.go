package main

import (
	"fmt"
)

func findSubstring(haystack, needle string) int {
	// if needle == "" {
	// 	return 0
	// }
	// return strings.Index(haystack, needle)

	// 1. Handle the edge case for an empty needle.
	if needle == "" {
		return 0
	}

	// Get the lengths one time.
	haystackLen := len(haystack)
	needleLen := len(needle)

	// 2. Loop through the haystack.
	// We only need to loop up to the point where the remaining
	// string is as long as the needle.
	for i := 0; i <= haystackLen-needleLen; i++ {
		// 3. Check if the slice of the haystack from i to i + needleLen
		// is equal to the needle.
		if haystack[i:i+needleLen] == needle {
			// If it matches, return the starting index i.
			return i
		}
	}

	// 4. If the loop finishes without a match, return -1.
	return -1
}

func main() {
	// Sample code in your main function:
	fmt.Println(findSubstring("hello world", "world"))
	fmt.Println(findSubstring("go is awesome", "is"))
	fmt.Println(findSubstring("foobar", "baz"))
	fmt.Println(findSubstring("testing", ""))
}
