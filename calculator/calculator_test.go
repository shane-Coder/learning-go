package calculator

import "testing"

func TestAdd(t *testing.T) {
	// 1. Define the test case struct
	testCases := []struct {
		name     string // A name for the sub-test
		a, b     int
		expected int
	}{
		// 2. Create the slice of test cases (our table)
		{name: "Positive Numbers", a: 5, b: 10, expected: 15},
		{name: "Negative Numbers", a: -5, b: -10, expected: -15},
		{name: "Mixed Numbers", a: 5, b: -10, expected: -5},
		{name: "Zero Value", a: 0, b: 0, expected: 0},
	}

	// 3. Loop over the test cases
	for _, tc := range testCases {
		// t.Run creates a sub-test, which gives us clean output
		t.Run(tc.name, func(t *testing.T) {
			// Call the function we are testing
			result := Add(tc.a, tc.b)

			// Check if the result matches the expectation
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
