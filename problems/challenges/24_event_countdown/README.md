# Challenge #24: Event Countdown Calculator

## Objective
Write a Go program with a function that calculates the remaining time until a future event. The function should take a date string, parse it, and return the duration until that date in a human-readable format (days, hours, minutes).

## Difficulty
Medium

## Concepts Tested
* Time Parsing (`time.Parse`)
* Time Duration and Difference (`time.Sub`, `time.Until`)
* Structs (`struct`)
* Integer and Floating-Point Math

## Rules/Logic
1.  Define a struct named `Countdown` with three fields: `Days` (int), `Hours` (int), and `Minutes` (int).
2.  Create a function `calculateCountdown(targetDate string) (Countdown, error)`. The input `targetDate` will be in the format `"YYYY-MM-DD HH:MM:SS"` (e.g., `"2025-12-25 00:00:00"`).
3.  Inside the function:
    * Define the layout string that matches the input date format.
    * Use `time.Parse` to convert the `targetDate` string into a `time.Time` object. Handle any parsing errors.
    * Get the current time using `time.Now()`.
    * Calculate the `time.Duration` remaining until the target date. A simple way is `target.Sub(now)`.
    * Check if the target date is in the past. If it is, return an error.
4.  Convert the total duration into days, hours, and minutes.
    * Hint: `duration.Hours()` returns the total duration in hours (e.g., 50.5 hours). You can use this to calculate the number of full days and the remaining hours.
5.  Populate and return the `Countdown` struct.
6.  In your `main` function, call your `calculateCountdown` with a future date and print the result in a formatted way. Also, test the error case with a date that has already passed.

## Your Tasks
1.  Create a new directory: `problems/challenges/24_event_countdown/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
// NOTE: The output will change depending on when you run the code.
// This example assumes the current time is 2025-10-05 10:00:00.

target := "2025-10-07 12:30:00"
countdown, err := calculateCountdown(target)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Time until %s: %d days, %d hours, %d minutes\n", target, countdown.Days, countdown.Hours, countdown.Minutes)
}

pastTarget := "2024-01-01 00:00:00"
_, err = calculateCountdown(pastTarget)
if err != nil {
    fmt.Println("Error:", err)
}

// Expected Terminal Output (assuming current time is 2025-10-05 10:00:00):
Time until 2025-10-07 12:30:00: 2 days, 2 hours, 30 minutes
Error: target date is in the past