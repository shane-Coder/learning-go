package main

import (
	"fmt"
	"sync"
)

// DatabaseConnection is the struct we want to make a singleton.
type DatabaseConnection struct {
	connectionString string
}

// --- Global variables for the Singleton Pattern ---

// 'instance' will hold our single connection object. It starts as nil.
var instance *DatabaseConnection

// 'once' is our special gatekeeper from the sync package.
var once sync.Once

// getInstance is the only way to get the singleton instance.
func getInstance() *DatabaseConnection {
	// The magic happens here!
	// once.Do() takes a function as an argument.
	// The first time getInstance() is called, this function will run, creating the instance.
	// On every subsequent call (from any goroutine), once.Do() will see that the function
	// has already been run and will do nothing.
	once.Do(func() {
		fmt.Println("Creating the database connection instance now.")
		instance = &DatabaseConnection{connectionString: "connected_to_db"}
	})

	// Return the single, shared instance.
	return instance
}

func main() {
	// A channel to collect the pointers from all our goroutines.
	results := make(chan *DatabaseConnection, 100)

	// Launch 100 goroutines concurrently.
	for i := 0; i < 100; i++ {
		go func() {
			// Each goroutine calls getInstance() to get the connection.
			conn := getInstance()
			// Send the pointer it received to the results channel.
			results <- conn
		}()
	}

	// --- Verification Step ---
	// Receive the first instance from the channel to use as our reference.
	firstInstance := <-results

	// Loop 99 more times to get the rest of the instances.
	for i := 0; i < 99; i++ {
		nextInstance := <-results
		// Compare the memory address of the next instance with our first one.
		// If they are not the same, the singleton failed.
		if firstInstance != nextInstance {
			fmt.Println("Error: Multiple instances were created!")
			return
		}
	}

	// If the loop completes, it means all 100 pointers were identical.
	fmt.Println("All 100 goroutines received the same singleton instance.")
	fmt.Printf("The instance's connection string is: %s\n", firstInstance.connectionString)
}
