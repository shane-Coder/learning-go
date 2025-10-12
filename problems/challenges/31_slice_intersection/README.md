# Challenge #31: Slice Intersection

## Objective
Write an efficient Go program to find the intersection of two integer slices. The intersection of two slices is a new slice containing only the elements that are present in **both** of the original slices. The result should contain no duplicates.

## Difficulty
Medium

## Concepts Tested
* Maps (`map[int]struct{}`)
* Slices and Loops
* Functions

## Rules/Logic
1.  Create a function `intersection(slice1, slice2 []int) []int`.
2.  The function should return a new slice containing only the unique elements that appear in both `slice1` and `slice2`.
3.  **The Challenge**: You should aim for an efficient solution. A naive approach would be to use nested loops, but this is slow (O(n*m) complexity). A much more efficient approach is to use a map (or a Set) as a lookup table.
4.  **Hint (Efficient Approach)**:
    * Create a Set (using a `map[int]struct{}`) from the first slice to store all its unique elements.
    * Create a second Set to store the elements that are in both slices.
    * Iterate through the second slice. For each element, check if it exists in the first set.
    * If it does, add that element to your intersection set. This automatically handles duplicates.
    * Finally, convert the intersection set back into a slice.
5.  In your `main` function, create two sample integer slices and call your `intersection` function to find and print their common elements.

## Your Tasks
1.  Create a new directory: `problems/challenges/31_slice_intersection/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
sliceA := []int{1, 2, 3, 4, 5}
sliceB := []int{4, 5, 6, 7, 8}

result := intersection(sliceA, sliceB)
fmt.Println("Intersection:", result)

sliceC := []int{1, 2, 2, 3, 4, 4, 5}
sliceD := []int{2, 4, 4, 6, 8}

result2 := intersection(sliceC, sliceD)
fmt.Println("Intersection with duplicates:", result2)

// Expected Terminal Output (order may vary):
Intersection: [4 5]
Intersection with duplicates: [2 4]