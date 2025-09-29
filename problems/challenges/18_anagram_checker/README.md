# Challenge #18: String Anagram Checker

## Objective
Write a Go program with a function that checks if two strings are anagrams of each other. Anagrams are words or phrases formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once (e.g., "listen" is an anagram of "silent").

## Difficulty
Medium

## Concepts Tested
* String Manipulation (`strings`)
* Runes and Sorting (`sort`)
* Maps (`map[rune]int`)
* Conditionals (`if/else`)

## Rules/Logic
1.  Create a function `areAnagrams(s1, s2 string) bool`.
2.  Inside the function, you must first "normalize" both strings. This means your comparison should be case-insensitive and should ignore any spaces or non-letter characters. For example, "A decimal point" and "I'm a dot in place." should be considered anagrams.
3.  The function should then determine if the two normalized strings are anagrams. There are two common ways to solve this:
    * **Sort and Compare**: Convert both normalized strings to rune slices, sort them, and then check if the resulting strings are identical.
    * **Frequency Map**: Create a frequency map (e.g., `map[rune]int`) of the characters in the first normalized string. Then, iterate through the second string, decrementing the counts in the map. If all counts are zero at the end, they are anagrams.
4.  Your function should return `true` if they are anagrams and `false` otherwise.
5.  In your `main` function, test your `areAnagrams` function with several pairs of strings, including simple cases, cases with different casing and spaces, and cases that are not anagrams.

## Your Tasks
1.  Create a new directory: `problems/challenges/18_anagram_checker/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
fmt.Println(areAnagrams("listen", "silent"))
fmt.Println(areAnagrams("hello", "world"))
fmt.Println(areAnagrams("Dormitory", "dirty room##"))
fmt.Println(areAnagrams("Go", "go go"))

// Expected Terminal Output:
true
false
true
false