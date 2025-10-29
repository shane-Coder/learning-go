# Challenge #48: Concurrent Website Status Checker

## Objective
Write a Go program that takes a slice of website URLs and checks their status (whether they are "UP" or "DOWN") concurrently, printing the results as they come in.

## Difficulty
Medium-Hard

## Concepts Tested
* Goroutines (`go`)
* Channels (`chan`)
* HTTP Client (`net/http`)
* Error Handling

## Rules/Logic
1.  Create a function `checkWebsite(url string, resultsChan chan<- string)`. This function will run as a goroutine.
2.  Inside `checkWebsite`:
    * Make an HTTP `GET` request to the given `url` using `http.Get()`.
    * If the `err` is not `nil`, send a formatted error message to the `resultsChan` (e.g., `fmt.Sprintf("%s is DOWN: %v", url, err)`).
    * If `err` is `nil`, it means the request was successful. **Don't forget to `defer resp.Body.Close()`**.
    * Send a success message to the `resultsChan` (e.g., `fmt.Sprintf("%s is UP (Status: %s)", url, resp.Status)`).
3.  In your `main` function:
    * Create a slice of website URLs to check. Include a mix of valid URLs (e.g., "https://www.google.com") and at least one invalid or non-existent URL (e.g., "https://this-site-does-not-exist.foo").
    * Create a channel for the results (e.g., `make(chan string)`).
    * Loop through your slice of URLs and, for each one, launch a `checkWebsite` goroutine.
    * After launching all the goroutines, use a **second loop** to receive the results from the channel. This loop must run exactly as many times as the number of websites you are checking.
    * Print each result as it comes in from the channel.

## Your Tasks
1.  Create a new directory: `problems/challenges/48_concurrent_website_checker/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
websites := []string{
    "[https://www.google.com](https://www.google.com)",
    "[https://www.github.com](https://www.github.com)",
    "[https://this-site-does-not-exist.foo](https://this-site-does-not-exist.foo)",
    "[https://www.bing.com](https://www.bing.com)",
}

resultsChan := make(chan string)

for _, url := range websites {
    go checkWebsite(url, resultsChan)
}

fmt.Println("--- Website Status ---")
for i := 0; i < len(websites); i++ {
    fmt.Println(<-resultsChan)
}
fmt.Println("----------------------")

// Expected Terminal Output (the order will vary due to concurrency):
--- Website Status ---
[https://www.google.com](https://www.google.com) is UP (Status: 200 OK)
[https://www.bing.com](https://www.bing.com) is UP (Status: 200 OK)
[https://this-site-does-not-exist.foo](https://this-site-does-not-exist.foo) is DOWN: Get "[https://this-site-does-not-exist.foo](https://this-site-does-not-exist.foo)": dial tcp: lookup this-site-does-not-exist.foo: no such host
[https://www.github.com](https://www.github.com) is UP (Status: 200 OK)
----------------------