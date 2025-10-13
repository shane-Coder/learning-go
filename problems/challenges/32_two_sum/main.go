package main

import (
	"errors"
	"fmt"
)

func twoSum(nums []int, target int) ([]int, error) {
	var store = make(map[int]int)
	for index, value := range nums {
		complement := target - value
		// Check if the complement exists in our map.
		if foundIndex, ok := store[complement]; ok {
			return []int{foundIndex, index}, nil
		}
		store[value] = index
	}
	return nil, errors.New("no two sum solution found")
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	indices1, err1 := twoSum(nums1, target1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Indices:", indices1) // Expected: [0, 1] or [1, 0]
	}

	nums2 := []int{3, 2, 4}
	target2 := 6
	indices2, err2 := twoSum(nums2, target2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Indices:", indices2) // Expected: [1, 2] or [2, 1]
	}

	nums3 := []int{3, 3}
	target3 := 6
	indices3, err3 := twoSum(nums3, target3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Indices:", indices3) // Expected: [0, 1] or [1, 0]
	}

	nums4 := []int{1, 2, 3, 4, 5}
	target4 := 10
	indices4, err4 := twoSum(nums4, target4)
	if err4 != nil {
		fmt.Println("Error:", err4) // Expected: Error: no two sum solution found
	} else {
		fmt.Println("Indices:", indices4)
	}
}
