# Challenge #1: Word Frequency Counter

## Objective
Write a Go program with a function that takes a block of text as a string and returns a map counting the frequency of each word.

## Difficulty
Easy

## Concepts Tested
* Maps (`map[string]int`)
* Slices (`[]string`)
* Standard Library (`strings` package)
* Functions

## Rules/Logic
1.  Create a function `countWords(text string) map[string]int`.
2.  Inside the function, you should first "normalize" the text to ensure the count is accurate:
    * Convert the entire text to lowercase using `strings.ToLower()`.
    * Remove common punctuation like periods (`.`) and commas (`,`) using `strings.ReplaceAll()`.
3.  Split the normalized text into a slice of words. `strings.Fields()` is a great tool for this as it handles multiple spaces correctly.
4.  Create an empty map to store the word counts.
5.  Loop over your slice of words. For each word, increment its count in the map.
6.  Return the final map.

## Your Tasks
1.  Create a new directory for this challenge: `problems/challenges/01_word_frequency/`.
2.  Inside this new directory, create two files: `README.md` and `main.go`.
3.  Copy and paste this full problem statement into your `README.md` file.
4.  Write your solution in the `main.go` file.
5.  In your `main` function, test your `countWords` function with a sample text and print the resulting map.
6.  Add your terminal output to the `README.md` and push your final solution to GitHub.

## Sample Input
```go
text := "Hello world, this is a test. Hello world again!"
Word frequencies: map[a:1 again:1 hello:2 is:1 test:1 this:1 world:2]