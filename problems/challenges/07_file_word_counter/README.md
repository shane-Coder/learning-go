# Challenge #7: File Word Counter

## Objective
Write a Go program that reads a text file, counts the frequency of each word, and prints a summary of the top 5 most common words. This is a classic data processing task.

## Difficulty
Medium

## Concepts Tested
* File I/O (`os`, `bufio`)
* Maps (`map[string]int`)
* String Manipulation (`strings`)
* Slices and Sorting (`sort`)
* Structs (`struct`)

## Rules/Logic
1.  Create a function `countWordsInFile(filePath string) (map[string]int, error)` that:
    * Opens the file specified by `filePath`.
    * Uses a `bufio.Scanner` to read the file line by line.
    * For each line, normalize it by converting to lowercase and removing common punctuation (like periods and commas).
    * Splits the line into words and updates a word frequency map.
    * Returns the final map and any error encountered.
2.  In your `main` function, you will first need a text file to read. You can create a file named `input.txt` with some sample text.
3.  Call your `countWordsInFile` function and handle any potential errors.
4.  To find the top 5 words, you'll need to sort them. Since you can't sort a map directly, the common Go pattern is:
    * Create a struct to hold a word and its count (e.g., `type WordCount struct { Word string; Count int }`).
    * Create a slice of this struct from your map.
    * Use the `sort.Slice` function to sort the slice in descending order of count.
5.  Finally, print the top 5 words from your sorted slice.

## Your Tasks
1.  Create a new directory: `problems/challenges/07_file_word_counter/`.
2.  Inside this directory, create three files: `README.md`, `main.go`, and `input.txt`.
3.  Copy and paste this problem statement into `README.md`.
4.  Add some sample text to `input.txt`. Use a few words multiple times.
5.  Write your solution in `main.go`.

## Sample `input.txt`
```text
Go is an open-source programming language.
Go is expressive, concise, clean, and efficient.
Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

--- Top 5 Most Common Words ---
1. and: 4
2. go: 3
3. language: 3
4. of: 3
5. the: 3