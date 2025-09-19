package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type LogEntry struct {
	Level   string
	Message string
}

func parseLogFile(filePath string) ([]LogEntry, error) {
	content, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer content.Close()
	var logEntries []LogEntry
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.SplitN(line, ":", 2)
		// Handle malformed lines. If there aren't 2 parts, skip the line.
		if len(text) != 2 {
			fmt.Printf("Warning: Skipping malformed log line: \"%s\"\n", line)
			continue
		}
		logEntries = append(logEntries, LogEntry{Level: strings.TrimSpace(text[0]), Message: strings.TrimSpace(text[1])})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return logEntries, nil
}

func main() {
	filePath, err := parseLogFile("problems/challenges/08_log_file_parser/logs.txt")
	if err != nil {
		log.Fatalf("Error parsing log file: %v", err)
	}
	fmt.Printf("Successfully parsed %d log entries.\n", len(filePath))
	fmt.Println("--- Log Summary ---")
	for _, line := range filePath {
		fmt.Printf("Level: %s, Message: %s\n", line.Level, line.Message)
	}
	fmt.Printf("Process done!!\n")
}
