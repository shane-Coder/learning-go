# Challenge #28: Sparse Matrix

## Objective
Write a Go program to implement a **Sparse Matrix**. A sparse matrix is a matrix in which most of the elements are zero. Instead of storing all the zeros in a 2D slice, it's much more memory-efficient to only store the non-zero elements. We will use a map of maps to achieve this.

## Difficulty
Hard

## Concepts Tested
* Maps of Maps (`map[int]map[int]int`)
* Structs and Methods
* Pointers
* Error Handling

## Rules/Logic
1.  Define a struct named `SparseMatrix` with two fields for dimensions (`Rows`, `Cols` int) and one field for the data: `elements` (a map of maps, `map[int]map[int]int`). The outer map's key is the row index, and the inner map's key is the column index, with the value being the non-zero element.
2.  Create a `NewSparseMatrix(rows, cols int) *SparseMatrix` constructor. It should initialize the `elements` map.
3.  Implement the following methods for the `SparseMatrix`.
    * `Set(row, col, value int) error`: Sets the value at a given `row` and `col`.
        * It must perform **bounds checking**. If the `row` or `col` is out of the matrix's dimensions, it should return an error.
        * If the `value` is `0`, you should *remove* the element from the map to save space. If the inner map for that row becomes empty, you can also remove the row's entry from the outer map.
        * If the `value` is non-zero, you must add it to the map. You will need to check if the inner map for that row exists and create it if it doesn't.
    * `Get(row, col int) (int, error)`: Gets the value at a given `row` and `col`.
        * It must also perform bounds checking and return an error if the indices are invalid.
        * If the element is found in the map, return its value.
        * If the element is **not** in the map, it means the value is `0`, so you should return `0, nil`.

## Your Tasks
1.  Create a new directory: `problems/challenges/28_sparse_matrix/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.
5.  In your `main` function, create a new `SparseMatrix`, set a few values, get a few values (including a zero value and an out-of-bounds value), and print the results to demonstrate it's working.

## Sample Input/Output

```go
// Sample code in your main function:
matrix := NewSparseMatrix(10, 10)

matrix.Set(0, 0, 10)
matrix.Set(2, 3, 20)
matrix.Set(5, 7, 30)
matrix.Set(2, 3, 0) // Set a value back to zero

val1, _ := matrix.Get(0, 0)
fmt.Println("Value at (0,0):", val1)

val2, _ := matrix.Get(5, 7)
fmt.Println("Value at (5,7):", val2)

val3, _ := matrix.Get(2, 3) // This was set to 0
fmt.Println("Value at (2,3):", val3)

_, err := matrix.Get(15, 15) // Out of bounds
if err != nil {
    fmt.Println("Error:", err)
}

// Expected Terminal Output:
Value at (0,0): 10
Value at (5,7): 30
Value at (2,3): 0
Error: index out of bounds