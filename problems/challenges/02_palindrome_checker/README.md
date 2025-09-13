# Challenge #2: Palindrome Checker

## Objective
Write a Go program with a function that checks if a given string is a palindrome. A palindrome is a word that reads the same forwards and backward, like "madam" or "racecar".

## Difficulty
Easy

## Concepts Tested
* Strings (`string`)
* Loops (`for`)
* Functions
* Conditionals (`if/else`)

## Rules/Logic
1.  Create a function `isPalindrome(text string) bool`.
2.  Inside the function, you should first "normalize" the input string to ensure an accurate check. This means ignoring spaces and case.
    * Remove all spaces from the string.
    * Convert the entire string to lowercase.
3.  The function should then check if the normalized string is the same forwards and backward.
4.  Return `true` if it's a palindrome and `false` if it is not. A simple way to do this is to reverse the string and compare it to the original.

## Your Tasks
1.  Create a new directory for this challenge: `problems/challenges/02_palindrome_checker/`.
2.  Inside this new directory, create two files: `README.md` and `main.go`.
3.  Copy and paste this full problem statement into your `README.md` file.
4.  Write your solution in the `main.go` file.
5.  In your `main` function, test your `isPalindrome` function with a few examples (like "madam", "hello", and "Taco Cat") and print the results.

## Sample Input/Output
```go
// Sample Input
fmt.Println(isPalindrome("madam"))
fmt.Println(isPalindrome("hello"))
fmt.Println(isPalindrome("Taco Cat")) // Should be true after normalizing
fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Should be true after all replacements of all non-alphanumeric characters

// Expected Terminal Output
true
false
true
true