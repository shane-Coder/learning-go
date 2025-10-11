# Challenge #30: Find the Missing Number

## Objective
Write an efficient Go program to find the single missing number in a slice of integers. The slice will contain `n` distinct numbers taken from the sequence `0, 1, 2, ..., n`.

## Difficulty
Medium-Hard

## Concepts Tested
* Algorithmic Thinking
* Mathematical Formulas (Gaussian Summation)
* Slices and Loops

## Rules/Logic
1.  Create a function `findMissingNumber(nums []int) int`.
2.  The input `nums` is a slice of integers. It is guaranteed to contain `n` unique numbers, and those numbers are all within the range `[0, n]`. This means exactly one number from that range is missing.
3.  **The Challenge**: You should aim for a solution that is more efficient than just sorting the slice or using a map.
4.  **Hint**: A very efficient way to solve this is using a mathematical formula. The sum of a sequence of numbers from `0` to `n` can be calculated with the formula: `expectedSum = n * (n + 1) / 2`.
    * You can calculate the `expectedSum`.
    * You can calculate the `actualSum` of the numbers in your input slice.
    * The difference between these two sums will be the missing number.
5.  In your `main` function, test your `findMissingNumber` function with a few different slices.

## Your Tasks
1.  Create a new directory: `problems/challenges/30_missing_number/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
nums1 := []int{3, 0, 1} // n = 3, range is [0, 3], missing number is 2
fmt.Println("Missing number is:", findMissingNumber(nums1))

nums2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1} // n = 9, range is [0, 9], missing is 8
fmt.Println("Missing number is:", findMissingNumber(nums2))

nums3 := []int{0, 1} // n = 2, range is [0, 2], missing is 2
fmt.Println("Missing number is:", findMissingNumber(nums3))

// Expected Terminal Output:
Missing number is: 2
Missing number is: 8
Missing number is: 2