package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Create a file. If it already exists, its content will be erased.
	file, err := os.Create("lessons/23_file_io_writing/output.txt")
	// 2. Immediately check for an error.
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return // Exit if we can't create the file
	}
	// 3. Defer the closing of the file.
	defer file.Close()

	// 4. Write a string to the file.
	content := "Hello from Go!\nThis is a new line.\nEnd of the file.\n"
	_, err = file.WriteString(content) // WriteString returns the number of bytes written and an error

	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}

	fmt.Println("File 'output.txt' created and written to successfully.")
}
