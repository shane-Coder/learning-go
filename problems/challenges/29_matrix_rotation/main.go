package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	// Step 1: Transpose the matrix.
	// We swap the element at (i, j) with the element at (j, i).
	// The inner loop starts from 'i' to avoid swapping elements back to their original place.
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row.
	// For each row, we use a two-pointer technique to reverse its elements.
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
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
}
