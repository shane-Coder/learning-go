package main

import (
	"errors"
	"fmt"
)

func binarySearch(sortedSlice []int, target int) (int, error) {
	left := 0
	right := len(sortedSlice) - 1
	for left <= right {
		mid := left + (right-left)/2
		if sortedSlice[mid] == target {
			return mid, nil
		} else if sortedSlice[mid] < target {
			left = mid + 1
		} else if sortedSlice[mid] > target {
			right = mid - 1
		}
	}
	return -1, errors.New("target not found in slice")
}

func main() {
	// Sample code in your main function:
	data := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	target1 := 23
	index1, err1 := binarySearch(data, target1)
	if err1 != nil {
		fmt.Printf("Error searching for %d: %v\n", target1, err1)
	} else {
		fmt.Printf("Found %d at index %d\n", target1, index1)
	}

	target2 := 40
	_, err2 := binarySearch(data, target2)
	if err2 != nil {
		fmt.Printf("Error searching for %d: %v\n", target2, err2)
	}
}
