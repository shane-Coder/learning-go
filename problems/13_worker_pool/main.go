/*
Problem Statement #13: Simple Worker Pool

Objective:
	Write a program that simulates a "worker" goroutine processing a "job". The main function will send a job (a simple string) to the worker via one channel and then wait for a confirmation of completion on a second channel.

Rules/Logic:

	1. Create two channels:

		- jobs chan string: To send work from main to the worker.
		- results chan string: To receive confirmation from the worker back in main.

	2. Create a worker function that takes both channels as arguments (worker(jobs <-chan string, results chan<- string)). Using <-chan and chan<- makes the channel directions explicit, which is good practice.

	3. Inside the worker function, use a for...range loop to receive jobs from the jobs channel.

	4. For each job received, print a message like "Worker processing job: [job]", simulate work with time.Sleep, and then send a completion message (e.g., "Finished job: [job]") to the results channel.

	5. In main:

		- Launch one worker goroutine.

		- Send one job (e.g., "Create a report") to the jobs channel.

		- Wait for and receive the confirmation from the results channel.

		- Print the confirmation.

---Terminal Output---
	// Testcase 1
	Main: Sending job 'Create a report'...
	Worker processing job: Create a report
	Finished job: Create a report
	Main: Received confirmation: 'Finished job: Create a report'

	// Testcase 2
	Main: Sending job 'Create a report'...
	Worker processing job: Create a report
	Finished job: Create a report
	Main: Received confirmation: 'Finished job: Create a report'

---
NOTE: This solution was developed with assistance to understand the core concepts
of channel communication and closing.
---
*/

package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan string, results chan<- string) {
	for job := range jobs {
		fmt.Printf("Worker processing job: %s\n", job)
		time.Sleep(5 * time.Second)
		completionMessage := fmt.Sprintf("Finished job: %s", job)
		fmt.Println(completionMessage)
		results <- completionMessage
	}
}

func main() {
	jobs := make(chan string, 1)
	results := make(chan string, 1)

	go worker(jobs, results)

	jobToSend := "Create a report"
	fmt.Printf("Main: Sending job '%s'...\n", jobToSend)
	jobs <- jobToSend

	close(jobs)

	confirmation := <-results
	fmt.Printf("Main: Received confirmation: '%s'\n", confirmation)
}
