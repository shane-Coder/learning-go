# Challenge #49: Simple Substring Search

## Objective
Write a Go program to implement a function that finds the first occurrence of a substring (the "needle") within a larger string (the "haystack"). You should **not** use the built-in `strings.Contains` or `strings.Index` functions.

## Difficulty
Medium

## Concepts Tested
* String, Rune, and Byte Slices
* Nested Loops
* Algorithmic Thinking

## Rules/Logic
1.  Create a function `findSubstring(haystack, needle string) int`.
2.  The function should search for the `needle` string within the `haystack` string.
3.  If the `needle` is found, the function should return the **starting index** of its first occurrence.
4.  If the `needle` is not found, the function should return `-1`.
5.  If the `needle` is an empty string, the function should return `0` (as per the standard Go library's convention).
6.  **Hint**:
    * You can iterate through the `haystack` with a `for` loop.
    * In each iteration, check if the slice of the `haystack` starting at your current index matches the `needle`.

## Your Tasks
1.  Create a new directory: `problems/challenges/49_substring_search/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
fmt.Println(findSubstring("hello world", "world"))
fmt.Println(findSubstring("go is awesome", "is"))
fmt.Println(findSubstring("foobar", "baz"))
fmt.Println(findSubstring("testing", ""))

// Expected Terminal Output:
6
3
-1
0