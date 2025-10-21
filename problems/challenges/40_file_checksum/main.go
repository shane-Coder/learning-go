package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

// calculateSHA256 takes a file path and returns its SHA-256 checksum as a hex string.
func calculateSHA256(filePath string) (string, error) {
	// 1. Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	// 2. CRITICAL: Defer the closing of the file. This runs after the function returns.
	defer file.Close()

	// 3. Create a new SHA-256 hasher.
	hasher := sha256.New()

	// 4. Copy the file's contents into the hasher.
	// io.Copy is very efficient for this. It takes a destination (Writer) and a source (Reader).
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	// 5. Get the final hash sum as a byte slice. The 'nil' argument is required.
	hashInBytes := hasher.Sum(nil)

	// 6. Convert the byte slice to a hexadecimal string.
	hashString := fmt.Sprintf("%x", hashInBytes)

	return hashString, nil
}

func main() {
	// 1. Check for the command-line argument.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	// 2. Call our function to calculate the checksum.
	checksum, err := calculateSHA256(filePath)
	if err != nil {
		// Use log.Fatalf for critical errors that should stop the program.
		log.Fatalf("Error calculating checksum: %v", err)
	}

	// 3. Print the final result.
	fmt.Printf("SHA-256 Checksum of '%s': %s\n", filePath, checksum)
}
