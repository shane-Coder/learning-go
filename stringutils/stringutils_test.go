/*
Problem Statement #17: Palindrome Checker Test

Objective:
	First, create a new function in our stringutils package to check if a string is a palindrome. Second, write a comprehensive, table-driven unit test to verify its correctness with multiple edge cases.

Rules/Logic:

	1. The Function (stringutils/stringutils.go):

		- Add a new, exported function to your stringutils package: IsPalindrome(s string) bool.

		- This function should return true if the string s is a palindrome and false otherwise.

		- A good implementation should be case-insensitive (e.g., "Madam" is a palindrome) and should ignore spaces (e.g., "taco cat" is a palindrome).

	2. The Test (stringutils/stringutils_test.go):

		- Create a new test file in the stringutils package named stringutils_test.go.

		- Inside this file, create a table-driven test function named TestIsPalindrome.

		- Your test table should include at least the following cases:

			A simple palindrome (e.g., "madam")

			A non-palindrome (e.g., "hello")

			A palindrome with mixed case (e.g., "Madam")

			A palindrome with spaces (e.g., "taco cat")

			An empty string ("")

		- Use t.Run() to create a sub-test for each case in your table.

---Terminal Output---
=== RUN   TestPalindrome
=== RUN   TestPalindrome/A_simple_palindrome
=== RUN   TestPalindrome/A_non-palindrome
=== RUN   TestPalindrome/A_palindrome_with_mixed_case
=== RUN   TestPalindrome/A_palindrome_with_spaces
=== RUN   TestPalindrome/An_empty_string
=== RUN   TestPalindrome/A_single_character
=== RUN   TestPalindrome/A_palindrome_with_spaces_and_mixed_case
--- PASS: TestPalindrome (0.00s)
	--- PASS: TestPalindrome/A_simple_palindrome (0.00s)
	--- PASS: TestPalindrome/A_non-palindrome (0.00s)
	--- PASS: TestPalindrome/A_palindrome_with_mixed_case (0.00s)
	--- PASS: TestPalindrome/A_palindrome_with_spaces (0.00s)
	--- PASS: TestPalindrome/An_empty_string (0.00s)
	--- PASS: TestPalindrome/A_single_character (0.00s)
	--- PASS: TestPalindrome/A_palindrome_with_spaces_and_mixed_case (0.00s)
PASS
ok  	learning-go/stringutils	0.123s
*/

package stringutils

import "testing"

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		information string
		palString   string
		expected    bool
		palindrome  string
	}{
		{information: "A simple palindrome", palString: "madam", expected: true},
		{information: "A non-palindrome", palString: "hello", expected: false},
		{information: "A palindrome with mixed case", palString: "Madam", expected: true},
		{information: "A palindrome with spaces", palString: "taco cat", expected: true},
		{information: "An empty string", palString: "", expected: true},
		{information: "A single character", palString: "a", expected: true},
		{information: "A palindrome with spaces and mixed case", palString: "A fox    a", expected: false},
	}
	for _, tc := range testCases {
		t.Run(tc.information, func(t *testing.T) {
			result, palindrome := IsPalindrome(tc.palString)
			if result != tc.expected {
				t.Errorf("string %s is not palindrome as %v and result string is %s", tc.palString, tc.expected, palindrome)
			}
		})
	}
}
