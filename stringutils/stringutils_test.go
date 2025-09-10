package stringutils

import "testing"

// This is our new test helper function.
func assertPalindrome(t *testing.T, input string, expected bool) {
	// t.Helper() marks this function as a test helper.
	// This is crucial for getting correct line numbers on failures.
	t.Helper()

	result, processedStr := IsPalindrome(input)
	if result != expected {
		t.Errorf("IsPalindrome(\"%s\") = %v; expected %v (processed as: '%s')", input, result, expected, processedStr)
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool // 'want' is a common name for the expected value
	}{
		{name: "Simple palindrome", input: "madam", want: true},
		{name: "Non-palindrome", input: "hello", want: false},
		{name: "Mixed case palindrome", input: "Madam", want: true},
		{name: "Palindrome with spaces", input: "taco cat", want: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// We now call our clean helper function.
			assertPalindrome(t, tc.input, tc.want)
		})
	}
}
