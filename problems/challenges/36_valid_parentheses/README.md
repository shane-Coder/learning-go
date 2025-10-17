# Challenge #36: Valid Parentheses

## Objective
Write a Go program to determine if a string containing just the characters '(', ')', '{', '}', '[' and ']' is valid. A string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Difficulty
Medium

## Concepts Tested
* Stacks (LIFO data structure)
* Maps (for bracket matching)
* Strings and Runes
* Conditionals

## Rules/Logic
1.  Create a function `isValid(s string) bool`.
2.  The function should return `true` if the input string `s` is valid, and `false` otherwise.
3.  **Hint**: The best way to solve this is using a **Stack**.
    * Iterate through the input string character by character (rune).
    * If you encounter an **opening bracket** ('(', '{', '['), push it onto the stack.
    * If you encounter a **closing bracket** (')', '}', ']'):
        * Check if the stack is empty. If it is, the string is invalid (no matching open bracket).
        * Pop the top element from the stack.
        * Check if the popped element is the correct corresponding opening bracket for the current closing bracket. If not, the string is invalid.
4.  After iterating through the entire string, the stack must be **empty** for the string to be valid. If there are any opening brackets left on the stack, it means they were never closed.
5.  **Optional**: A map can be useful to quickly find the matching opening bracket for a closing bracket.

## Your Tasks
1.  Create a new directory: `problems/challenges/36_valid_parentheses/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`. You can reuse the `Stack` implementation from Challenge #14, or implement a simpler one directly in this file if you prefer.

## Sample Input/Output

```go
// Sample code in your main function:
fmt.Println(isValid("()"))      // true
fmt.Println(isValid("()[]{}"))  // true
fmt.Println(isValid("(]"))      // false
fmt.Println(isValid("([)]"))    // false
fmt.Println(isValid("{[]}"))    // true
fmt.Println(isValid(""))        // true
fmt.Println(isValid("["))        // false
fmt.Println(isValid("]"))        // false

// Expected Terminal Output:
true
true
false
false
true
true
false
false