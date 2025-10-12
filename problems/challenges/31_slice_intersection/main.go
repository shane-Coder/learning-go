package main

import "fmt"

func intersection(slice1, slice2 []int) []int {
	set1 := make(map[int]struct{})
	for _, val := range slice1 {
		set1[val] = struct{}{}
	}
	intersectionSet := make(map[int]struct{})
	for _, val := range slice2 {
		if _, ok := set1[val]; ok {
			intersectionSet[val] = struct{}{}
		}
	}
	result := make([]int, 0, len(intersectionSet))
	for val := range intersectionSet {
		result = append(result, val)
	}

	return result
}

func main() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{4, 5, 6, 7, 8}

	result := intersection(sliceA, sliceB)
	fmt.Println("Intersection:", result)

	sliceC := []int{1, 2, 2, 3, 4, 4, 5}
	sliceD := []int{2, 4, 4, 6, 8}

	result2 := intersection(sliceC, sliceD)
	fmt.Println("Intersection with duplicates:", result2)
}
