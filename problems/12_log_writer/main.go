/*
Problem Statement #12: Simple Log Writer

Objective:
	Write a Go program that appends a new log entry with a timestamp to a file named log.txt each time it's run.

Rules/Logic:

	1. Create a function logMessage(message string) error.

	2. Inside this function, open the file log.txt. You must open it in append mode. (Hint: Use os.OpenFile with the flags os.O_APPEND|os.O_CREATE|os.O_WRONLY).

	3. Ensure you defer file.Close().

	4. Get the current time using time.Now() and format it into a string.

	5. Create the full log line, including the timestamp and the message (e.g., 2025-09-05 11:43:00 - This is a log message.).

	6. Write the new log line to the file.

	7. In your main function, call logMessage with a sample message and handle any potential errors.

---Terminal Output---
	// Test case 1
	$ go run main.go
	Log entry added successfully.

	// Test case 2
	$ go run main.go
	Log entry added successfully.

---log.txt---
	2025-09-05 12:05:17 - Application started.
	2025-09-05 12:06:21 - Application started successfully.
	2025-09-05 12:06:30 - Application started successfully.
	2025-09-05 12:09:15 - This is a sample log message.
	2025-09-05 12:14:25 - This is a sample log message.
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func logMessage(message string) error {
	// Open the file in append mode, create it if it doesn't exist, and set write-only permissions
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the current time and format it
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Create the full log line
	logLine := fmt.Sprintf("%s - %s\n", currentTime, message)

	// Write the log line to the file
	if _, err := file.WriteString(logLine); err != nil {
		return err
	}

	return nil
}

func main() {
	// Call logMessage with a sample message
	if err := logMessage("This is a sample log message."); err != nil {
		fmt.Printf("Error logging message: %v\n", err)
	} else {
		fmt.Println("Log entry added successfully.")
	}
}
