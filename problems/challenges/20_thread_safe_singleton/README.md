# Challenge #20: Thread-Safe Singleton

## Objective
Write a Go program to implement a **thread-safe Singleton pattern**. The Singleton pattern ensures that a class (or in Go, a struct) has only one instance and provides a global point of access to it. Your implementation must be safe to use from multiple goroutines concurrently.

## Difficulty
Hard

## Concepts Tested
* Design Patterns (Singleton)
* Concurrency (`sync.Once`, Goroutines)
* Pointers
* Structs
* Channels (for collecting results)

## Rules/Logic
1.  Define a struct that you want to be a singleton. Let's call it `DatabaseConnection`. It can have a simple field, like `connectionString` (string).
2.  You will need a global variable to hold the single instance of your struct and another global variable of type `sync.Once`. The `sync.Once` type is a special tool in Go that guarantees a piece of code will run **exactly one time**, no matter how many goroutines try to run it.
3.  Create a function `getInstance() *DatabaseConnection`. This function is the only way to get an instance of your `DatabaseConnection`.
4.  Inside `getInstance()`, you must use `once.Do()`. The function you pass to `once.Do()` will contain the logic to create the `DatabaseConnection` instance. Because of how `sync.Once` works, this creation logic will only ever be executed once, even if a thousand goroutines call `getInstance()` at the exact same time.
5.  In your `main` function, you need to prove that it works.
    * Launch 100 goroutines.
    * Inside each goroutine, call `getInstance()` and send the returned pointer to a channel.
    * In `main`, receive all 100 pointers from the channel.
    * Check if all the pointers are identical. You can do this by comparing the first pointer you receive to all the subsequent ones. If they are all the same, your Singleton is working correctly.

## Your Tasks
1.  Create a new directory: `problems/challenges/20_thread_safe_singleton/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Expected Terminal Output
```text
All 100 goroutines received the same singleton instance.
The instance's connection string is: connected_to_db