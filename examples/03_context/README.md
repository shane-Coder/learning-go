# Go Context Package: A Deep Dive

## Objective
This example provides a practical guide to Go's `context` package, one of the most important packages for building robust, concurrent, and networked applications. It demonstrates how to handle cancellations and timeouts for long-running operations.

---

## What is a `Context`?

A `context.Context` is an object that carries signals across function calls in a concurrent program. Think of it as a "lifeline" for a request or task. It's primarily used for:

1.  **Cancellation**: Telling a function to stop its work because the result is no longer needed.
2.  **Timeouts/Deadlines**: Automatically cancelling an operation if it takes too long.
3.  **Passing Request-Scoped Data**: Carrying data like a user's ID or a request trace ID through a chain of function calls.

---

## How It Works

The core of the `context` package is the `Done()` method, which returns a channel. This channel is **closed** when the context is cancelled. We can use a `select` statement to listen on this channel and stop our work gracefully.

### Example 1: `context.WithTimeout`

This is used to create a context that will be automatically cancelled after a specified duration. It's perfect for ensuring that operations like database queries or API calls don't hang forever.

### Example 2: `context.WithCancel`

This is used to create a context that we can cancel manually. It returns a `cancel` function. When we call this function, the context's `Done()` channel is closed. This is useful when a user action (like closing a browser tab) should trigger the cancellation of a background task.

---

## Running the Example
To see both methods in action, navigate to this directory and run:
```bash
go run main.go