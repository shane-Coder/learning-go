# Challenge #33: Find the First Non-Repeating Character

## Objective
Write an efficient Go program to find the first character in a string that does not repeat. For example, in the string "swiss", the first non-repeating character is "w".

## Difficulty
Medium

## Concepts Tested
* Maps for Frequency Counting (`map[rune]int`)
* String and Rune Iteration (`for...range`)
* Algorithmic Thinking

## Rules/Logic
1.  Create a function `firstNonRepeating(s string) (rune, error)`.
2.  The function should take a string as input.
3.  It needs to find the first character (from left to right) that appears only once in the entire string.
4.  If a non-repeating character is found, the function should return that character (`rune`) and a `nil` error.
5.  If all characters in the string repeat, it should return a zero value for the rune and an error (e.g., `errors.New("no non-repeating character found")`).
6.  **The Challenge**: An efficient solution involves two passes over the string (or data derived from it), not nested loops.
7.  **Hint (Efficient Approach)**:
    * **First Pass**: Iterate through the string and build a frequency map (`map[rune]int`) that counts how many times each character appears.
    * **Second Pass**: Iterate through the string *again*, from beginning to end. For each character, look up its count in your frequency map. The very first character you find that has a count of `1` is your answer.

## Your Tasks
1.  Create a new directory: `problems/challenges/33_first_non_repeating_char/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, map-based solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
s1 := "swiss"
char1, err1 := firstNonRepeating(s1)
if err1 != nil {
    fmt.Println("Error:", err1)
} else {
    fmt.Printf("First non-repeating char in '%s' is '%c'\n", s1, char1)
}

s2 := "aabbcc"
_, err2 := firstNonRepeating(s2)
if err2 != nil {
    fmt.Println("Error:", err2)
}

// Expected Terminal Output:
First non-repeating char in 'swiss' is 'w'
Error: no non-repeating character found