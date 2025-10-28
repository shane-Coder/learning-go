package main

import (
	"fmt"
	"time"
)

func rateLimitedPrinter(rate time.Duration, messages []string) {
	// 1. Create the ticker with the specified rate.
	ticker := time.NewTicker(rate)
	// 2. Ensure the ticker is stopped when the function exits.
	defer ticker.Stop()

	// 3. Loop through the slice of message strings.
	for _, msg := range messages {
		// 4. Wait here until the ticker sends a signal on its channel.
		<-ticker.C
		// 5. Once the signal is received, print the current message.
		fmt.Println(msg)
	}
}

func main() {
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

	messagesToPrint = []string{
		"Another Message 1",
		"Another Message 2",
		"Another Message 3",
	}
	printRate = 1 * time.Second // Print one message every second

	fmt.Println("Starting another rate-limited printing...")
	rateLimitedPrinter(printRate, messagesToPrint)
	fmt.Println("Finished.")
}
