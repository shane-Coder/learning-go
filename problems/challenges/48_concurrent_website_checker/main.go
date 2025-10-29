package main

import (
	"fmt"
	"net/http" // We need net/http to make requests
	"sync"     // We'll use a WaitGroup for a cleaner exit, though the loop in main is also fine.
)

// checkWebsite performs the HTTP GET request and sends the result to the channel.
func checkWebsite(url string, resultsChan chan<- string, wg *sync.WaitGroup) {
	// Signal to the WaitGroup that this goroutine is done when the function returns.
	defer wg.Done()

	// 1. Capture the (response, error) from http.Get.
	resp, err := http.Get(url)

	// 2. Handle the error case FIRST.
	if err != nil {
		// 3. Send the error message to the results channel.
		resultsChan <- fmt.Sprintf("%s is DOWN: %v", url, err)
		return // Exit the function early.
	}

	// 4. If there was no error, we MUST close the response body.
	defer resp.Body.Close()

	// 5. Send the success message to the results channel.
	resultsChan <- fmt.Sprintf("%s is UP (Status: %s)", url, resp.Status)
}

func main() {
	// 6. Fix the URLs - they should be clean strings.
	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://this-site-does-not-exist.foo",
		"https://www.bing.com",
	}

	resultsChan := make(chan string)
	var wg sync.WaitGroup

	// --- Goroutine Launch ---
	for _, url := range websites {
		// Increment the WaitGroup counter for each goroutine.
		wg.Add(1)
		// Launch the worker goroutine.
		go checkWebsite(url, resultsChan, &wg)
	}

	// --- Result Collection (Alternative Professional Pattern) ---
	// Launch a separate goroutine to close the channel *after* all workers are done.
	// This lets us use a clean for...range loop below.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	fmt.Println("--- Website Status ---")
	// This for...range loop will automatically stop when the channel is closed.
	for result := range resultsChan {
		fmt.Println(result)
	}
	fmt.Println("----------------------")

	// // Sample code in your main function:
	// websites := []string{
	// 	"[https://www.google.com](https://www.google.com)",
	// 	"[https://www.github.com](https://www.github.com)",
	// 	"[https://this-site-does-not-exist.foo](https://this-site-does-not-exist.foo)",
	// 	"[https://www.bing.com](https://www.bing.com)",
	// }

	// resultsChan := make(chan string)

	// for _, url := range websites {
	// 	go checkWebsite(url, resultsChan)
	// }

	// fmt.Println("--- Website Status ---")
	// for i := 0; i < len(websites); i++ {
	// 	fmt.Println(<-resultsChan)
	// }
	// fmt.Println("----------------------")
}
