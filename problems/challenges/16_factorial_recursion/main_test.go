// main_test.go
package main

import "testing"

func TestFactorial(t *testing.T) {
	// 1. Define the "table" - a slice of test case structs.
	testCases := []struct {
		name          string // A descriptive name for the test case
		input         int
		expectedValue int
		expectError   bool // Do we expect an error?
	}{
		// 2. Add the individual test cases to the table.
		{
			name:          "Factorial of 5",
			input:         5,
			expectedValue: 120,
			expectError:   false,
		},
		{
			name:          "Base case: Factorial of 0",
			input:         0,
			expectedValue: 1,
			expectError:   false,
		},
		{
			name:          "Error case: Negative number",
			input:         -2,
			expectedValue: 0,
			expectError:   true,
		},
	}

	// 3. Loop over the table of test cases.
	for _, tc := range testCases {
		// t.Run creates a sub-test, which gives us clean output for each case.
		t.Run(tc.name, func(t *testing.T) {
			// Run the function we are testing.
			result, err := factorial(tc.input)

			// 4. Check the results.
			if tc.expectError {
				// If we expected an error, but didn't get one...
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				// If we did NOT expect an error, but got one...
				if err != nil {
					t.Errorf("did not expect an error but got one: %v", err)
				}
				// If the result is not what we expected...
				if result != tc.expectedValue {
					t.Errorf("factorial(%d) = %d; expected %d", tc.input, result, tc.expectedValue)
				}
			}
		})
	}
}
