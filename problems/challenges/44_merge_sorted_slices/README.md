# Challenge #44: Merge Two Sorted Slices

## Objective
Write an efficient Go program to merge two already sorted integer slices into a single sorted slice.

## Difficulty
Medium

## Concepts Tested
* Slices and Indices
* Looping (`for` or `while`-style `for`)
* Pointers/Indices Management
* Algorithmic Thinking

## Rules/Logic
1.  Create a function `mergeSortedSlices(slice1, slice2 []int) []int`.
2.  The input `slice1` and `slice2` are **guaranteed** to be already sorted in ascending order.
3.  The function should return a **new** slice that contains all elements from both `slice1` and `slice2`, also sorted in ascending order.
4.  **The Challenge**: You should aim for an efficient solution that runs in **O(n + m)** time complexity, where n and m are the lengths of the two input slices. This means you should only need to iterate through each slice once.
5.  **Hint (Efficient Approach - Two Pointers)**:
    * Create a new result slice with a capacity large enough to hold all elements (`len(slice1) + len(slice2)`).
    * Initialize three index pointers: `i` for `slice1`, `j` for `slice2`, and `k` for the `result` slice, all starting at 0.
    * Use a loop that continues as long as both `i` and `j` are within the bounds of their respective slices (`i < len(slice1)` and `j < len(slice2)`).
    * Inside the loop, compare `slice1[i]` and `slice2[j]`. Append the smaller element to the `result` slice and increment the corresponding pointer (`i` or `j`).
    * After the main loop finishes, one of the slices might still have elements left over. Append all remaining elements from that slice to the `result`.

## Your Tasks
1.  Create a new directory: `problems/challenges/44_merge_sorted_slices/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, two-pointer solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
list1 := []int{1, 3, 5, 7}
list2 := []int{2, 4, 6, 8}
merged1 := mergeSortedSlices(list1, list2)
fmt.Println("Merged:", merged1)

list3 := []int{10, 20}
list4 := []int{5, 15, 25}
merged2 := mergeSortedSlices(list3, list4)
fmt.Println("Merged:", merged2)

// Expected Terminal Output:
Merged: [1 2 3 4 5 6 7 8]
Merged: [5 10 15 20 25]