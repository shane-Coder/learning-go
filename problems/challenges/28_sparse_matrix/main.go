package main

import (
	"errors"
	"fmt"
)

// SparseMatrix stores non-zero elements in a map of maps.
type SparseMatrix struct {
	Rows     int
	Cols     int
	elements map[int]map[int]int
}

// NewSparseMatrix is a constructor that initializes and returns a new SparseMatrix.
func NewSparseMatrix(rows, cols int) *SparseMatrix {
	return &SparseMatrix{
		Rows:     rows,
		Cols:     cols,
		elements: make(map[int]map[int]int), // Initialize the outer map.
	}
}

// Set sets the value at a given row and col.
func (sm *SparseMatrix) Set(row, col, value int) error {
	// 1. Bounds Checking: Ensure the row and col are within the matrix dimensions.
	if row < 0 || row >= sm.Rows || col < 0 || col >= sm.Cols {
		return errors.New("index out of bounds")
	}

	// 2. Handle Zero Values: If the value is 0, we remove the element to save space.
	if value == 0 {
		// Check if the row exists in the outer map.
		if rowMap, ok := sm.elements[row]; ok {
			// If it exists, delete the element for that column.
			delete(rowMap, col)
			// If the row's map is now empty, delete the row itself.
			if len(rowMap) == 0 {
				delete(sm.elements, row)
			}
		}
		return nil
	}

	// 3. Handle Non-Zero Values.
	// Check if the inner map for this row exists.
	if _, ok := sm.elements[row]; !ok {
		// If it doesn't exist, create it.
		sm.elements[row] = make(map[int]int)
	}
	// Now that we know the inner map exists, set the value.
	sm.elements[row][col] = value
	return nil
}

// Get retrieves the value at a given row and col.
func (sm *SparseMatrix) Get(row, col int) (int, error) {
	// 1. Bounds Checking.
	if row < 0 || row >= sm.Rows || col < 0 || col >= sm.Cols {
		return 0, errors.New("index out of bounds")
	}

	// 2. Look up the value.
	// Check if the row exists.
	if rowMap, ok := sm.elements[row]; ok {
		// If the row exists, check if the column exists in the inner map.
		if value, ok := rowMap[col]; ok {
			return value, nil
		}
	}

	// 3. If the element is not found in the map, it means its value is 0.
	return 0, nil
}

func main() {
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
}
