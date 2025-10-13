# Challenge #32: Two Sum Problem

## Objective
Write an efficient Go program to solve the "Two Sum" problem. Given a slice of integers and a target value, you need to find the indices of the two numbers in the slice that add up to the target.

## Difficulty
Medium-Hard

## Concepts Tested
* Algorithmic Thinking
* Maps for Lookups (`map[int]int`)
* Slices and Loops
* Error Handling

## Rules/Logic
1.  Create a function `twoSum(nums []int, target int) ([]int, error)`.
2.  The function should take a slice of integers and a target integer.
3.  It needs to find two numbers in the slice that sum up to the `target`.
4.  If a pair is found, the function should return a slice containing the **indices** of the two numbers.
5.  If no such pair is found, it should return `nil` and an error.
6.  You can assume that each input will have exactly one solution, and you may not use the same element twice.
7.  **The Challenge**: A simple approach is to use nested loops to check every pair of numbers, but this is slow (O(n²)). A much more efficient **O(n)** solution uses a map to store the numbers you've already seen and their indices.
8.  **Hint (Efficient Approach)**:
    * Create an empty map to store numbers and their indices (`map[int]int`).
    * Iterate through the input slice `nums` with both the index and the value.
    * For each number, calculate the `complement` needed to reach the target (i.e., `complement = target - number`).
    * Check if the `complement` already exists in your map.
        * If it does, you have found your pair! Return the index of the complement (which you stored in the map) and the index of the current number.
        * If it doesn't, add the current number and its index to the map, then continue to the next iteration.

## Your Tasks
1.  Create a new directory: `problems/challenges/32_two_sum/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, map-based solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
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

// Expected Terminal Output:
Indices: [0 1]
Indices: [1 2]
Indices: [0 1]