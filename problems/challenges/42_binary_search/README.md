# Challenge #42: Binary Search

## Objective
Write a Go program to implement the **Binary Search** algorithm to find the index of a target value within a **sorted** slice of integers.

## Difficulty
Medium-Hard

## Concepts Tested
* Algorithmic Thinking (Divide and Conquer)
* Slices and Indices
* Loops (`for`)
* Conditionals (`if/else`)

## Rules/Logic
1.  Create a function `binarySearch(sortedSlice []int, target int) (int, error)`.
2.  The input `sortedSlice` is **guaranteed** to be sorted in ascending order.
3.  The function should search for the `target` integer within the `sortedSlice`.
4.  **Binary Search Algorithm**:
    * Initialize two pointers, `low` at the beginning of the slice (index 0) and `high` at the end of the slice (index `len(slice)-1`).
    * While `low` is less than or equal to `high`:
        * Calculate the `mid` index: `mid = low + (high-low)/2` (this prevents potential integer overflow compared to `(low+high)/2`).
        * Compare the element at `sortedSlice[mid]` with the `target`:
            * If `sortedSlice[mid] == target`, you've found the target! Return `mid, nil`.
            * If `sortedSlice[mid] < target`, the target must be in the right half of the current search space. Update `low = mid + 1`.
            * If `sortedSlice[mid] > target`, the target must be in the left half. Update `high = mid - 1`.
5.  If the loop finishes without finding the target (meaning `low` has become greater than `high`), the target is not in the slice. Return `-1` and an error (e.g., `errors.New("target not found in slice")`).
6.  In your `main` function, create a sorted slice of integers and test your `binarySearch` function with targets that exist in the slice and targets that do not.

## Your Tasks
1.  Create a new directory: `problems/challenges/42_binary_search/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
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

// Expected Terminal Output:
Found 23 at index 5
Error searching for 40: target not found in slice