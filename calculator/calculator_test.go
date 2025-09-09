package calculator

import "testing"

func TestAdd(t *testing.T) {
	// Define the test case
	a := 10
	b := 5
	expected := 15

	// Call the function we are testing
	result := Add(a, b)

	// Check if the result matches the expectation
	if result != expected {
		// If it doesn't match, report an error.
		// t.Errorf() will fail the test and print a formatted error message.
		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, result, expected)
	}
}
