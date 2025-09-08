package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Create an in-memory io.Reader from a simple string.
	// The strings.NewReader function does this for us.
	reader := strings.NewReader("Hello from the io.Reader!")

	// os.Stdout is a pre-existing variable that represents the console output.
	// It satisfies the io.Writer interface.
	writer := os.Stdout

	fmt.Println("--- Copying from a Reader to a Writer ---")

	// io.Copy will read all data from the reader and write it to the writer.
	// It returns the number of bytes copied and any error.
	bytesCopied, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("Error copying data:", err)
	} else {
		fmt.Printf("\nSuccessfully copied %d bytes.\n", bytesCopied)
	}
}
