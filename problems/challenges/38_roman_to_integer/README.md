# Challenge #38: Roman to Integer

## Objective
Write a Go program with a function that converts a Roman numeral string (e.g., "MCMXCIV") into its corresponding integer value (1994).

## Difficulty
Medium

## Concepts Tested
* Maps (`map[rune]int`)
* Strings and Runes
* Looping and Conditionals

## Rules/Logic
1.  Roman numerals are represented by seven different symbols: `I`, `V`, `X`, `L`, `C`, `D`, and `M`.
    * I = 1
    * V = 5
    * X = 10
    * L = 50
    * C = 100
    * D = 500
    * M = 1000
2.  Numerals are usually written largest to smallest from left to right (e.g., `XII` = 10 + 1 + 1 = 12).
3.  **Subtraction Rule**: There are six instances where subtraction is used:
    * `I` placed before `V` (5) or `X` (10) makes 4 (`IV`) and 9 (`IX`).
    * `X` placed before `L` (50) or `C` (100) makes 40 (`XL`) and 90 (`XC`).
    * `C` placed before `D` (500) or `M` (1000) makes 400 (`CD`) and 900 (`CM`).
4.  Create a function `romanToInt(s string) int`.
5.  **Hint**: The easiest way to handle the subtraction rule is to iterate through the Roman numeral string from **left to right**. Keep track of the total value. For each numeral, look at the *next* numeral. If the next numeral is larger than the current one, subtract the current numeral's value from the total. Otherwise, add the current numeral's value to the total.
6.  A map can be useful to store the integer value of each Roman numeral character.

## Your Tasks
1.  Create a new directory: `problems/challenges/38_roman_to_integer/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
fmt.Println("III =", romanToInt("III"))
fmt.Println("LVIII =", romanToInt("LVIII"))
fmt.Println("MCMXCIV =", romanToInt("MCMXCIV"))

// Expected Terminal Output:
III = 3
LVIII = 58
MCMXCIV = 1994