# Challenge #29: Matrix Rotation

## Objective
Write a Go program with a function that rotates an `N x N` 2D matrix (represented as a slice of slices) by 90 degrees clockwise **in-place**. "In-place" means you should not create a new matrix; you must modify the existing one directly.

## Difficulty
Hard

## Concepts Tested
* 2D Slices (`[][]int`)
* Loops (nested `for` loops)
* In-place Algorithm Design

## Rules/Logic
1.  Create a function `rotate(matrix [][]int)`. The matrix is guaranteed to be square (`N x N`).
2.  The function must rotate the matrix 90 degrees clockwise. For example, the top row becomes the rightmost column, the rightmost column becomes the bottom row, and so on.
3.  The rotation must be done **in-place**, meaning you modify the `matrix` slice directly without creating a second 2D slice for the result.
4.  A common two-step approach to solve this is:
    * **Transpose the matrix**: Swap the element at `(row, col)` with the element at `(col, row)`.
    * **Reverse each row**: Iterate through each row and reverse its elements.
5.  In your `main` function, create a sample `N x N` matrix, print it, call your `rotate` function, and then print the matrix again to show the result.

## Your Tasks
1.  Create a new directory: `problems/challenges/29_matrix_rotation/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

fmt.Println("Original Matrix:")
// A helper function to print the matrix might be nice!
for _, row := range matrix {
    fmt.Println(row)
}

rotate(matrix)

fmt.Println("\nRotated Matrix:")
for _, row := range matrix {
    fmt.Println(row)
}

// Expected Terminal Output:
Original Matrix:
[1 2 3]
[4 5 6]
[7 8 9]

Rotated Matrix:
[7 4 1]
[8 5 2]
[9 6 3]