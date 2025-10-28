# Challenge #47: Rate Limiter Simulation

## Objective
Write a Go program that simulates a simple rate limiter using a **Ticker**. A rate limiter is essential for controlling how often an action (like an API request) can be performed within a given time window.

## Difficulty
Medium-Hard

## Concepts Tested
* Concurrency (`time.Ticker`, Channels)
* Goroutines
* The `time` Package
* Loops

## Rules/Logic
1.  Create a function `rateLimitedPrinter(rate time.Duration, messages []string)`.
    * `rate`: The allowed time duration between each message (e.g., `500 * time.Millisecond`).
    * `messages`: A slice of strings to be printed.
2.  Inside the function:
    * Create a new **Ticker** using `time.NewTicker(rate)`. This ticker will send a value on its channel (`ticker.C`) at the specified interval.
    * **Important**: Remember to `defer ticker.Stop()` to release the ticker's resources when the function finishes.
    * Use a `for...range` loop to iterate through the `messages`.
    * Inside the loop, **wait** for the ticker to send a signal by receiving from its channel (`<-ticker.C`). This will block the loop, ensuring that you only proceed at the specified `rate`.
    * After receiving the signal, print the current message.
3.  In your `main` function:
    * Create a slice of sample messages.
    * Define a rate (e.g., half a second).
    * Call your `rateLimitedPrinter` function.

## Your Tasks
1.  Create a new directory: `problems/challenges/47_rate_limiter/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
messagesToPrint := []string{
    "Message 1",
    "Message 2",
    "Message 3",
    "Message 4",
    "Message 5",
}
printRate := 500 * time.Millisecond // Print one message every half second

fmt.Println("Starting rate-limited printing...")
rateLimitedPrinter(printRate, messagesToPrint)
fmt.Println("Finished.")

// Expected Terminal Output (with a half-second pause between each message):
Starting rate-limited printing...
Message 1
Message 2
Message 3
Message 4
Message 5
Finished.