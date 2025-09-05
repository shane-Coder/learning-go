package main

import (
	"fmt"
	"time"
)

// This function will run concurrently as a goroutine.
func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

func main() {
	// Launch the sayHello function as a new goroutine.
	// The program does NOT wait for it to finish.
	go sayHello()

	// The main function continues its own execution immediately.
	fmt.Println("Hello from the main function!")

	// We sleep for a moment to give the goroutine a chance to run.
	// Without this, the main function might exit before the goroutine starts.
	time.Sleep(1 * time.Second)
}
