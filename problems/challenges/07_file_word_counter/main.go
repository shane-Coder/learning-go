package main

import (
	"bufio"
	"fmt"
	"log" // Using the log package is great for fatal errors
	"os"
	"sort"
	"strings"
)

// countWordsInFile reads a file and returns a frequency map of its words.
func countWordsInFile(filePath string) (map[string]int, error) {
	// 1. Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// Return nil for the map and the error if opening fails
		return nil, err
	}
	// 2. CRITICAL: Defer the closing of the file.
	defer file.Close()

	wordMap := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// 3. Loop through each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		// 4. Normalize the line
		line = strings.ToLower(line)
		// A better way to remove punctuation without adding extra spaces
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ".", "")

		// 5. Split the line into words and update the map
		words := strings.Fields(line)
		for _, word := range words {
			wordMap[word]++
		}
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 6. Return the map and a nil error for success
	return wordMap, nil
}

// WordCount is a struct to hold a word and its frequency, for sorting.
type WordCount struct {
	Word  string
	Count int
}

func main() {
	// Call the function to get the word map
	wordMap, err := countWordsInFile("problems/challenges/07_file_word_counter/input.txt")
	if err != nil {
		// If there's an error, log it and exit the program
		log.Fatalf("Error reading file: %v", err)
	}

	// --- Convert Map to Slice for Sorting ---
	// 1. Create an empty slice of our WordCount struct.
	var wordCounts []WordCount
	// 2. Loop over the map and add each word/count to the slice.
	for word, count := range wordMap {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	// --- Sort the Slice ---
	// 3. Use sort.Slice with a custom "less" function.
	// We want to sort in DESCENDING order, so we check if the count at index i is GREATER than at j.
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	// --- Print the Top 5 ---
	fmt.Println("--- Top 5 Most Common Words ---")
	// Add a check to prevent crashing if there are fewer than 5 words
	numWordsToShow := 5
	if len(wordCounts) < numWordsToShow {
		numWordsToShow = len(wordCounts)
	}

	for i := 0; i < numWordsToShow; i++ {
		fmt.Printf("%d. %s: %d\n", i+1, wordCounts[i].Word, wordCounts[i].Count)
	}
}
