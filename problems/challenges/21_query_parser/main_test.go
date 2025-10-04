package main

import (
	"reflect" // We use the 'reflect' package for deep map comparison
	"testing"
)

func TestParseQuery(t *testing.T) {
	// 1. Define the "table" of test cases.
	testCases := []struct {
		name        string
		input       string
		expectedMap map[string]string
		expectError bool
	}{
		// 2. Add the individual test cases.
		{
			name:  "Valid query string",
			input: "name=shivam&city=gurgaon",
			expectedMap: map[string]string{
				"name": "shivam",
				"city": "gurgaon",
			},
			expectError: false,
		},
		{
			name:        "Malformed parameter with no equals sign",
			input:       "name=shivam&malformed",
			expectedMap: nil,
			expectError: true,
		},
		{
			name:        "Malformed parameter with empty key",
			input:       "=shivam&city=gurgaon",
			expectedMap: nil,
			expectError: true,
		},
	}

	// 3. Loop over the table.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function we are testing.
			resultMap, err := parseQuery(tc.input)

			// 4. Check for errors.
			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got one: %v", err)
				}
				// 5. Check if the resulting map is correct.
				// reflect.DeepEqual is the standard way to compare complex types like maps.
				if !reflect.DeepEqual(resultMap, tc.expectedMap) {
					t.Errorf("got %v; want %v", resultMap, tc.expectedMap)
				}
			}
		})
	}
}
