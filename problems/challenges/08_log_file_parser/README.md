# Challenge #8: Log File Parser

## Objective
Write a Go program that reads a structured log file, parses each line into a struct, and then prints a formatted summary of the log entries.

## Difficulty
Medium

## Concepts Tested
* File I/O (`os`, `bufio`)
* Structs (`struct`)
* String Manipulation (`strings.SplitN`)
* Error Handling (`if err != nil`)
* Slices of Structs (`[]LogEntry`)

## Rules/Logic
1.  Each line in the log file will have a specific format: `LEVEL: MESSAGE`. For example: `INFO: User logged in`.
2.  Define a struct named `LogEntry` with two fields: `Level` (string) and `Message` (string).
3.  Create a function `parseLogFile(filePath string) ([]LogEntry, error)` that:
    * Opens and reads the specified log file line by line.
    * For each line, it splits the line into the `LEVEL` and `MESSAGE` parts. Use `strings.SplitN(line, ":", 2)` to safely split only on the first colon.
    * It should handle potential malformed lines (e.g., lines without a colon) gracefully, perhaps by printing a warning and skipping them.
    * It creates a `LogEntry` struct for each valid line and appends it to a slice.
    * It returns the slice of `LogEntry` structs and any error that occurred.
4.  In your `main` function:
    * Create a sample `logs.txt` file with a few log entries, including at least one malformed line.
    * Call your `parseLogFile` function.
    * If no error occurs, loop through the returned slice of `LogEntry` structs and print them in a clean, formatted way.

## Your Tasks
1.  Create a new directory: `problems/challenges/08_log_file_parser/`.
2.  Inside this directory, create `README.md`, `main.go`, and `logs.txt`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Add the sample log content to `logs.txt`.
5.  Write your solution in `main.go`.

## Sample `logs.txt`
```text
INFO: User logged in successfully.
WARNING: Disk space is running low.
ERROR: Database connection failed.
This is a malformed line.
INFO: User logged out.

Warning: Skipping malformed log line: "This is a malformed line."
Successfully parsed 4 log entries.
--- Log Summary ---
Level: INFO, Message: User logged in successfully.
Level: WARNING, Message: Disk space is running low.
Level: ERROR, Message: Database connection failed.
Level: INFO, Message: User logged out.