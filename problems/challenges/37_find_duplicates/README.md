# Challenge #37: Find All Duplicates in a Slice

## Objective
Write an efficient Go program to find all the numbers that appear more than once in a given slice of integers.

## Difficulty
Medium

## Concepts Tested
* Maps for Frequency Counting (`map[int]int`)
* Slices and Loops
* Functions

## Rules/Logic
1.  Create a function `findDuplicates(nums []int) []int`.
2.  The function should take a slice of integers as input.
3.  It needs to identify all numbers that appear **two or more times** in the slice.
4.  The function should return a new slice containing these duplicate numbers. Each duplicate number should appear only **once** in the result slice, regardless of how many times it appeared in the input. The order of numbers in the result slice doesn't matter.
5.  **Hint**: An efficient approach involves using a map to count the frequency of each number.
    * Iterate through the input slice and build a frequency map (`map[int]int`).
    * Iterate through the frequency map. Any number with a count greater than 1 is a duplicate. Add these numbers to your result slice.

## Your Tasks
1.  Create a new directory: `problems/challenges/37_find_duplicates/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, map-based solution in `main.go`.

## Sample Input/Output

```go
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


// Expected Terminal Output (order may vary):
Duplicates: [2 3]
Duplicates: [1]
Duplicates: []