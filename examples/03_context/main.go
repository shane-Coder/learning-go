// examples/03_context/main.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// slowOperation simulates a task that takes a long time, like a database query.
// It takes a context so it can be cancelled, and a WaitGroup to signal completion.
func slowOperation(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Slow operation started...")

	// The 'select' statement listens on multiple channels.
	// It will proceed with whichever case becomes ready first.
	select {
	case <-time.After(3 * time.Second):
		// This case happens if the 3-second timer finishes before the context is cancelled.
		fmt.Println("Slow operation finished successfully.")
	case <-ctx.Done():
		// This case happens if the context's Done() channel is closed.
		// This occurs when the context is cancelled or times out.
		// ctx.Err() provides the reason for the cancellation.
		fmt.Printf("Slow operation cancelled: %v\n", ctx.Err())
	}
}

func main() {
	var wg sync.WaitGroup

	// --- Example 1: Cancellation with a Timeout ---
	fmt.Println("--- Running example with a 1-second timeout ---")
	timeoutCtx, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelTimeout() // Good practice to call cancel to release resources.

	wg.Add(1)
	go slowOperation(timeoutCtx, &wg)
	wg.Wait() // Wait for the timeout example to finish.

	// --- Example 2: Manual Cancellation (The Practice Challenge) ---
	fmt.Println("\n--- Running example with manual cancellation ---")
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	wg.Add(1)
	go slowOperation(cancelCtx, &wg)

	// Let the goroutine run for a moment before we manually cancel it.
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Sending cancellation signal...")
	cancelFunc() // Manually trigger the cancellation.

	wg.Wait() // Wait for the manual cancel example to finish.
	fmt.Println("Main function finished.")
}
