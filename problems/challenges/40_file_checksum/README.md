# Challenge #40: File Checksum Calculator

## Objective
Write a Go program that calculates the SHA-256 checksum of a given file. A checksum is a unique signature for a file, and it's commonly used to verify that a file has not been corrupted during download or transfer.

## Difficulty
Medium

## Concepts Tested
* File I/O (`os.Open`, `io.Copy`)
* Cryptography (`crypto/sha256`)
* Formatted Printing (`fmt.Printf`)
* Command-Line Arguments (`os.Args`)

## Rules/Logic
1.  Your program should accept a file path as a command-line argument.
2.  Create a function `calculateSHA256(filePath string) (string, error)`.
3.  Inside this function:
    * Open the file specified by `filePath`. Handle any errors. Remember to `defer file.Close()`.
    * Create a new SHA-256 hash object using `sha256.New()`.
    * Copy the entire contents of the file into the hash object. `io.Copy()` is the perfect, most efficient tool for this. It takes a destination `Writer` (the hash object) and a source `Reader` (the file). Handle any errors from the copy.
    * Get the final hash sum as a byte slice from the hash object.
    * Convert the byte slice into a hexadecimal string using `fmt.Sprintf("%x", byteSlice)`.
4.  Return the hexadecimal string and a `nil` error on success.
5.  In your `main` function:
    * Check for the command-line argument. If it's missing, print a usage message.
    * Call your `calculateSHA256` function.
    * Print the resulting checksum or any error that occurred.

## Your Tasks
1.  Create a new directory: `problems/challenges/40_file_checksum/`.
2.  Inside this directory, create `README.md`, `main.go`, and a sample file `sample.txt` with some text in it.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## How to Test
1.  Run your program, passing the path to your `sample.txt` file as an argument.
    ```bash
    go run main.go sample.txt
    ```
2.  Change a single character in `sample.txt` and run the program again. You should see that the checksum is now completely different.

## Sample `sample.txt`
```text
Hello, Go World!

SHA-256 Checksum of 'sample.txt': 2b803a61003c10a19c5324508933b9347895b607a72661ab553f47c3e5907436