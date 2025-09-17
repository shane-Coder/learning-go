# Challenge #6: Concurrent Task Processor

## Objective
Write a Go program that simulates a pool of workers processing a list of jobs concurrently. This is a very common pattern for improving the performance of applications that need to handle many independent tasks.

## Difficulty
Medium

## Concepts Tested
* Goroutines (`go`)
* Channels (`chan`)
* WaitGroups (`sync.WaitGroup`)
* Structs (`struct`)
* The `time` Package

## Rules/Logic
1.  Define a struct named `Job` with two fields: `ID` (int) and `Duration` (a `time.Duration`, e.g., `2 * time.Second`).
2.  Create a `worker` function that takes two arguments: an `id` (int, for the worker's number) and a channel that sends jobs (`jobs <-chan Job`).
3.  The `worker` function should use a `for...range` loop to receive jobs from the channel.
4.  For each job, the worker should print a message that it's starting the job, then "work" by sleeping for the job's duration (`time.Sleep(job.Duration)`), and finally print a message that it has finished the job.
5.  In your `main` function:
    * Create a list (slice) of `Job` structs with varying durations.
    * Create a buffered channel for the jobs.
    * Launch 3 `worker` goroutines.
    * Send all the jobs from your list to the jobs channel.
    * After sending all the jobs, you must **close the channel**. This is a signal to the workers that no more jobs will be sent.
    * **Important**: You need a way to make the `main` function wait for all the goroutines to finish their work before it exits. A `sync.WaitGroup` is the perfect tool for this.

## Your Tasks
1.  Create a new directory: `problems/challenges/06_concurrent_task_processor/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md` file.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
jobs := []Job{
    {ID: 1, Duration: 2 * time.Second},
    {ID: 2, Duration: 1 * time.Second},
    {ID: 3, Duration: 3 * time.Second},
    {ID: 4, Duration: 1 * time.Second},
    {ID: 5, Duration: 2 * time.Second},
}

// Expected Terminal Output (the order may vary due to concurrency):
Worker 2 starting job 2
Worker 1 starting job 1
Worker 3 starting job 3
Worker 2 finished job 2
Worker 2 starting job 4
Worker 2 finished job 4
Worker 2 starting job 5
Worker 1 finished job 1
Worker 2 finished job 5
Worker 3 finished job 3
All jobs have been processed.