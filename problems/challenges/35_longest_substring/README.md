# Challenge #35: Longest Substring Without Repeating Characters

## Objective
Write an efficient Go program to find the length of the longest substring within a given string that does not contain any repeating characters.

## Difficulty
Hard

## Concepts Tested
* Algorithmic Thinking (Sliding Window Technique)
* Maps for Character Lookups (`map[rune]int`)
* Strings and Runes
* Math (`math.Max`)

## Rules/Logic
1.  Create a function `lengthOfLongestSubstring(s string) int`.
2.  The function should find the longest substring inside `s` that has no duplicate characters and return its length. For example, in `"abcabcbb"`, the longest substring without repeating characters is `"abc"`, so the length is 3.
3.  **The Challenge**: A naive solution with nested loops is too slow. The efficient solution uses a "sliding window" technique with a map.
4.  **Hint (Sliding Window Approach)**:
    * You will have a "window" represented by a `start` and `end` pointer (or the loop's index).
    * Use a map to keep track of the most recent index of each character you've seen (`map[rune]int`).
    * Iterate through the string with your `end` pointer (the `for...range` loop index).
    * For each character, check if it's already in your map.
        * If it is, it means you have a repeat. The start of your *new* window must be *after* the last place you saw that character. Update your `start` pointer accordingly.
    * After handling any repeats, update the character's current index in the map.
    * In each step, calculate the length of the current window (`end - start + 1`) and keep track of the maximum length you've seen so far.

## Your Tasks
1.  Create a new directory: `problems/challenges/35_longest_substring/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, map-based solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
s1 := "abcabcbb"
fmt.Printf("Length of longest substring in '%s': %d\n", s1, lengthOfLongestSubstring(s1))

s2 := "bbbbb"
fmt.Printf("Length of longest substring in '%s': %d\n", s2, lengthOfLongestSubstring(s2))

s3 := "pwwkew"
fmt.Printf("Length of longest substring in '%s': %d\n", s3, lengthOfLongestSubstring(s3))

// Expected Terminal Output:
Length of longest substring in 'abcabcbb': 3
Length of longest substring in 'bbbbb': 1
Length of longest substring in 'pwwkew': 3