# Challenge #16: Factorial Calculator (Recursion)

## Objective
Write a Go program that uses a **recursive function** to calculate the factorial of a non-negative integer. The factorial of a number `n` (written as `n!`) is the product of all positive integers less than or equal to `n`. For example, `5! = 5 * 4 * 3 * 2 * 1 = 120`.

## Difficulty
Medium

## Concepts Tested
* Recursion
* Functions
* Conditionals (`if/else`)
* Error Handling

## Rules/Logic
1.  Create a function `factorial(n int) (int, error)`.
2.  This function must be **recursive**. This means it will call itself with a smaller version of the problem until it reaches a base case.
3.  **Base Case**: The factorial of `0` is `1`. This is the condition that stops the recursion. If `n == 0`, the function should return `1, nil`.
4.  **Recursive Step**: For any number `n` greater than 0, the factorial is `n` multiplied by the factorial of `n-1`. The function should call itself with `n-1` and use the result.
5.  **Error Handling**: Factorial is not defined for negative numbers. If `n` is less than 0, the function should return `0` and an error (e.g., `errors.New("factorial is not defined for negative numbers")`).
6.  In your `main` function, test your `factorial` function with a few different numbers, including `5`, `0`, and a negative number to check your error handling.

## Your Tasks
1.  Create a new directory: `problems/challenges/16_factorial_recursion/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
result, err := factorial(5)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Factorial of 5 is:", result)
}

result, _ = factorial(0)
fmt.Println("Factorial of 0 is:", result)

_, err = factorial(-2)
if err != nil {
    fmt.Println("Error:", err)
}

// Expected Terminal Output:
Factorial of 5 is: 120
Factorial of 0 is: 1
Error: factorial is not defined for negative numbers