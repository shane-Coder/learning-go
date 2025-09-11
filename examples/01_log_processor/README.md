# Example #1: Concurrent Log Processor

## Objective
This example demonstrates a professional, real-world concurrency pattern in Go. The program reads log entries from a file, uses a pool of worker goroutines to parse them in parallel, and safely collects the results into a final slice.

This solution was developed collaboratively as a deep dive into advanced concurrency, including debugging a deadlock and implementing a race-free worker/collector pattern.

## Rules/Logic
1.  **The Input**: An input file named `logs.txt` with lines in the format `LEVEL:MESSAGE`.
2.  **The Struct**: A `LogEntry` struct to hold the parsed `Level` and `Message`.
3.  **The Workers**: A pool of concurrent `worker` goroutines to parse each line.
4.  **The Collector**: A dedicated `collector` goroutine to safely receive results from the workers.
5.  **The `main` Function**: Orchestrates the process by reading the file, dispatching lines to workers via channels, and using `WaitGroup` to ensure all goroutines finish before the program exits.

## Terminal Output
```text
--- All Parsed Logs ---
Level: INFO, Message: User logged in
Level: ERROR, Message: Database connection failed
Level: DEBUG, Message: Request received for /api/users
Level: INFO, Message: Processing payment