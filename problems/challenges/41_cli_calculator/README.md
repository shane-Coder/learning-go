# Challenge #41: Command-Line Calculator

## Objective
Write a Go program that acts as a simple calculator, taking three command-line arguments: two numbers and an operator. The program should perform the calculation and print the result.

## Difficulty
Medium

## Concepts Tested
* Command-Line Arguments (`os.Args`)
* String to Float Conversion (`strconv.ParseFloat`)
* Control Flow (`switch` statement)
* Error Handling

## Rules/Logic
1.  The program must accept exactly three command-line arguments: `number1`, `operator`, and `number2`.
2.  Create a function `calculate(num1 float64, operator string, num2 float64) (float64, error)`.
3.  The `calculate` function should use a `switch` statement on the `operator` to perform the correct operation. It should support addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
4.  The function must handle division by zero. If the operator is `/` and `num2` is `0`, it should return an error.
5.  If the operator is not one of the four supported types, it should also return an error.
6.  In your `main` function:
    * Check if the correct number of command-line arguments (`os.Args` should have a length of 4) is provided. If not, print a usage message and exit.
    * Parse the first and third arguments from strings to `float64` using `strconv.ParseFloat`. Handle any parsing errors.
    * The second argument is the operator string.
    * Call your `calculate` function with the parsed values.
    * Print the result or any error that occurred.

## Your Tasks
1.  Create a new directory: `problems/challenges/41_cli_calculator/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## How to Test
Run your program from the terminal with different arguments:

```bash
go run main.go 10 + 5
go run main.go 100 - 20.5
go run main.go 5 "*" 4
go run main.go 20 / 4
go run main.go 10 / 0
go run main.go 10 x 5

10.00 + 5.00 = 15.00
100.00 - 20.50 = 79.50
5.00 * 4.00 = 20.00
20.00 / 4.00 = 5.00
Error: division by zero is not allowed
Error: invalid operator: x