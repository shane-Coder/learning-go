package main

import "fmt"

func findDuplicates(nums []int) []int {
	count := make(map[int]int)
	result := []int{}
	for _, val := range nums {
		count[val] += 1
	}
	for key, value := range count {
		if value > 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	// Sample code in your main function:
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	duplicates1 := findDuplicates(nums1)
	fmt.Println("Duplicates:", duplicates1) // Order doesn't matter, e.g., [2 3] or [3 2]

	nums2 := []int{1, 1, 2}
	duplicates2 := findDuplicates(nums2)
	fmt.Println("Duplicates:", duplicates2) // Should contain only [1]

	nums3 := []int{1, 2, 3}
	duplicates3 := findDuplicates(nums3)
	fmt.Println("Duplicates:", duplicates3) // Should be empty []
}
