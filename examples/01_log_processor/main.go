package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type LogEntry struct {
	Level   string
	Message string
}

func worker(id int, lines <-chan string, results chan<- LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d started", id)

	for line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			entry := LogEntry{
				Level:   strings.TrimSpace(parts[0]),
				Message: strings.TrimSpace(parts[1]),
			}
			results <- entry
		} else {
			log.Printf("Worker %d: malformed line skipped: %s", id, line)
		}
	}
	log.Printf("Worker %d finished", id)
}

func main() {
	file, err := os.Open("logs.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	linesChan := make(chan string)
	resultsChan := make(chan LogEntry)
	var workerWg sync.WaitGroup

	numWorkers := 2
	workerWg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, linesChan, resultsChan, &workerWg)
	}

	var parsedLogs []LogEntry
	var collectorWg sync.WaitGroup
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for entry := range resultsChan {
			parsedLogs = append(parsedLogs, entry)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	close(linesChan)

	workerWg.Wait()
	close(resultsChan)
	collectorWg.Wait()

	log.Println("All processing complete.")
	fmt.Println("\n--- All Parsed Logs ---")
	for _, entry := range parsedLogs {
		fmt.Printf("Level: %s, Message: %s\n", entry.Level, entry.Message)
	}
}
