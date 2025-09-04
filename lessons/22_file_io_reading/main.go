package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- Attempting to read data.txt ---")

	// 1. Open the file.
	file, err := os.Open("data.txt")
	// 2. Immediately check for an error.
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return // Exit if we can't open the file
	}
	// 3. Defer the closing of the file. This will run when main() exits.
	defer file.Close()

	// 4. Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	lineNumber := 1
	// 5. Loop through each line of the file.
	for scanner.Scan() {
		// Get the current line of text.
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNumber, line)
		lineNumber++
	}

	// After the loop, check if the scanner encountered any errors.
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
}
