package main

import "fmt"

// flatten takes a nested structure (int or []interface{}) and returns a flat []int.
func flatten(nested interface{}) []int {
	// Initialize an empty slice to store the flattened integers.
	result := []int{}
	// Use a type switch to determine the actual type stored in the interface{}.
	switch val := nested.(type) {
	case int:
		// Base Case: If it's an integer, append it directly to the result.
		result = append(result, val)
	case []interface{}:
		for _, element := range val {
			// Recursive Step: If it's a slice of interfaces...
			flattenElement := flatten(element)
			// ...and append the *elements* of the returned slice to our result.
			result = append(result, flattenElement...)
		}
		// Default case (optional): Handle unexpected types if necessary.
		// default:
		//     fmt.Printf("Warning: encountered unexpected type %T\n", v)
	}
	return result
}

func main() {
	// Sample code in your main function:
	nestedList := []interface{}{
		1,
		[]interface{}{2, 3},
		4,
		[]interface{}{
			5,
			[]interface{}{6, 7},
		},
		8,
	}

	flattenedList := flatten(nestedList)
	fmt.Println("Flattened list:", flattenedList) // Expected: [1 2 3 4 5 6 7 8]

	// Test with a simpler case
	simpleList := []interface{}{10, 20, 30}
	flattenedSimple := flatten(simpleList)
	fmt.Println("Flattened simple list:", flattenedSimple) // Expected: [10 20 30]

	// Test with an empty slice
	emptyList := []interface{}{}
	flattenedEmpty := flatten(emptyList)
	fmt.Println("Flattened empty list:", flattenedEmpty) // Expected: []
}
