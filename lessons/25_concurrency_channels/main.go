package main

import (
	"fmt"
	"time"
)

// This worker function takes a channel as an argument.
// It will send a signal on this channel when its work is done.
func worker(done chan bool) {
	fmt.Println("Worker: Starting work...")
	// Simulate doing some work.
	time.Sleep(100 * time.Second)
	fmt.Println("Worker: Work is finished.")

	// Send a value (in this case 'true') on the channel to signal completion.
	done <- true
}

func main() {
	// Create a new channel that can transport boolean values.
	doneChannel := make(chan bool)

	fmt.Println("Main: Launching worker goroutine...")
	// Launch the worker concurrently.
	go worker(doneChannel)

	fmt.Println("Main: Waiting for worker to finish...")

	// This is a blocking receive. The main goroutine will pause here
	// until it receives a value from the doneChannel.
	<-doneChannel

	fmt.Println("Main: Worker has finished. Program can now exit.")
}
