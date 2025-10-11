package main

import "fmt"

func findMissingNumber(nums []int) int {
	// The array contains n distinct numbers in the range [0, n].
	// The length of the array is n, so one number in the range is missing.
	// We can use the formula for the sum of the first n natural numbers:
	// sum = n * (n + 1) / 2
	// Then subtract the sum of the elements in the array from this expected sum.
	// The difference will be the missing number.
	// Calculate the expected sum of numbers from 0 to n
	// where n is the length of the array
	// and the actual sum of the elements in the array.
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, val := range nums {
		actualSum += val
	}
	return expectedSum - actualSum
}

func main() {
	nums1 := []int{3, 0, 1} // n = 3, range is [0, 3], missing number is 2
	fmt.Println("Missing number is:", findMissingNumber(nums1))

	nums2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1} // n = 9, range is [0, 9], missing is 8
	fmt.Println("Missing number is:", findMissingNumber(nums2))

	nums3 := []int{0, 1} // n = 2, range is [0, 2], missing is 2
	fmt.Println("Missing number is:", findMissingNumber(nums3))
}
