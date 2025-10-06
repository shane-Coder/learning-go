# Challenge #25: Character Frequency Counter

## Objective
Write a Go program with a function that takes a string and returns a map counting the frequency of each character (rune). The program should handle multi-byte Unicode characters correctly and ignore spaces.

## Difficulty
Easy-Medium

## Concepts Tested
* Maps (`map[rune]int`)
* Runes
* String Iteration (`for...range`)
* Conditionals (`if`)

## Rules/Logic
1.  Create a function `countCharacters(text string) map[rune]int`.
2.  Inside the function, create an empty map that will hold the character counts. The map's keys should be `rune`s and its values should be `int`s.
3.  Loop through the input `text` string. The `for...range` loop is perfect for this as it automatically gives you one `rune` at a time, correctly handling Unicode.
4.  For each character:
    * If the character is a space, ignore it.
    * Otherwise, increment its count in your map.
5.  Return the final map.
6.  In your `main` function:
    * Create a sample string, including some spaces and repeated characters (and maybe a multi-byte character like an emoji).
    * Call your `countCharacters` function.
    * Loop through the resulting map and print each character and its frequency.

## Your Tasks
1.  Create a new directory: `problems/challenges/25_char_frequency_counter/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
text := "hello world 👋"
frequencyMap := countCharacters(text)

fmt.Println("--- Character Frequencies ---")
for char, count := range frequencyMap {
    fmt.Printf("'%c': %d\n", char, count)
}

// Expected Terminal Output (the order may vary):
--- Character Frequencies ---
'h': 1
'e': 1
'l': 3
'o': 2
'w': 1
'r': 1
'd': 1
'👋': 1