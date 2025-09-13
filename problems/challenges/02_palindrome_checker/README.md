# Challenge #2: Palindrome Checker (Upgraded)

## Objective
Write a Go program with an efficient function that checks if a given string is a palindrome, ignoring case, spaces, and all non-alphanumeric characters.

## Difficulty
Medium

## Concepts Tested
* String Manipulation
* Loops (`for`) & Conditionals (`if/else`)
* Unicode Package (`unicode`)
* Efficient String Building (`strings.Builder`)
* Algorithms (Two-Pointer Technique)

## Rules/Logic
1.  Create a function `isPalindrome(text string) bool`.
2.  Inside the function, "normalize" the input string by creating a new string that contains only the alphanumeric characters from the original, converted to lowercase.
    * Use `unicode.IsLetter` and `unicode.IsDigit` to identify the valid characters.
    * Use a `strings.Builder` for efficient string construction.
3.  Implement a **two-pointer technique** on the normalized string to check if it's a palindrome.
    * Set one pointer at the start of the string and one at the end.
    * Compare the characters at both pointers, moving them towards the center.
    * If a mismatch is found at any point, return `false` immediately.
4.  If the pointers meet or cross without finding a mismatch, the string is a palindrome. Return `true`. **Do not create a new reversed string.**

## Your Tasks
1.  Create a new directory for this challenge: `problems/challenges/02_palindrome_checker/`.
2.  Inside this new directory, create two files: `README.md` and `main.go`.
3.  Copy and paste this full problem statement into your `README.md` file.
4.  Write your efficient, upgraded solution in the `main.go` file.
5.  In your `main` function, test your `isPalindrome` function with several complex examples.

## Sample Input/Output
```go
// Sample Input
fmt.Println(isPalindrome("madam"))
fmt.Println(isPalindrome("hello"))
fmt.Println(isPalindrome("Taco Cat"))
fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
fmt.Println(isPalindrome("No 'x' 121 x on"))

// Expected Terminal Output
true
false
true
true
true