package main

import (
	"fmt"
)

func mergeSortedSlices(slice1, slice2 []int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < len(slice1) {
		for j < len(slice2) {
			if slice1[i] <= slice2[j] {
				result = append(result, slice1[i])
				i += 1
			} else {
				result = append(result, slice2[j])
				j += 1
			}
			if i == len(slice1) && j < len(slice2) {
				for j < len(slice2) {
					result = append(result, slice2[j])
					j += 1
				}
			}
			if i < len(slice1) && j == len(slice2) {
				for i < len(slice1) {
					result = append(result, slice1[i])
					i += 1
				}
			}
		}
	}
	return result
}

func main() {
	// Sample code in your main function:
	list1 := []int{1, 3, 5, 7}
	list2 := []int{2, 4, 6, 8}
	merged1 := mergeSortedSlices(list1, list2)
	fmt.Println("Merged:", merged1)

	list3 := []int{10, 20}
	list4 := []int{5, 15, 25}
	merged2 := mergeSortedSlices(list3, list4)
	fmt.Println("Merged:", merged2)
}
