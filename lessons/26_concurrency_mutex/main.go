package main

import (
	"fmt"
	"sync" // The 'sync' package contains the Mutex
)

var (
	counter = 0
	mutex   sync.Mutex     // Our mutex to protect the counter
	wg      sync.WaitGroup // A WaitGroup waits for a collection of goroutines to finish
)

func increment() {
	// defer wg.Done() ensures this is called when the function exits
	defer wg.Done()

	// Lock the mutex before accessing the counter.
	// This blocks any other goroutine from entering this section.
	mutex.Lock()
	// The code between Lock() and Unlock() is the critical section.
	counter++
	// Unlock the mutex so another goroutine can acquire it.
	mutex.Unlock()
}

func main() {
	// We want to wait for 1000 goroutines to finish.
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go increment()
	}

	// wg.Wait() blocks until all 1000 goroutines have called wg.Done().
	// This is a much better way to wait than time.Sleep()!
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
