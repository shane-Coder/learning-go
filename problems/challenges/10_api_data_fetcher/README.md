# Challenge #10: API Data Fetcher

## Objective
Write a Go program that fetches data from a public JSON API, parses the response into a Go struct, and prints a formatted summary. This is a fundamental task for any program that needs to interact with web services.

## Difficulty
Medium

## Concepts Tested
* Making HTTP GET Requests (`net/http`)
* Reading an HTTP Response Body (`io.ReadAll`)
* JSON Unmarshalling (`encoding/json`)
* Structs (`struct`)
* Error Handling (`if err != nil`)

## Rules/Logic
1.  We will use the **JSONPlaceholder API**, a free fake API for testing and prototyping. Specifically, the endpoint for a single "todo" item: `https://jsonplaceholder.typicode.com/todos/1`.
2.  Define a `Todo` struct to match the JSON response from the API. The JSON looks like this:
    ```json
    {
        "userId": 1,
        "id": 1,
        "title": "delectus aut autem",
        "completed": false
    }
    ```
    Remember to use struct tags (e.g., `` `json:"userId"` ``) to map the JSON keys to your Go struct fields.
3.  Create a function `fetchTodo(todoID int) (*Todo, error)` that:
    * Constructs the correct URL string for the given `todoID`.
    * Makes an HTTP `GET` request to that URL using `http.Get()`.
    * Reads the entire response body using `io.ReadAll()`.
    * Unmarshals the JSON body into a `Todo` struct.
    * Returns a pointer to the `Todo` struct and any error that occurred.
4.  In your `main` function:
    * Call your `fetchTodo` function with an ID (e.g., 1).
    * Handle any potential errors.
    * If the fetch is successful, print the details of the fetched to-do item in a clean, formatted way.

## Your Tasks
1.  Create a new directory: `problems/challenges/10_api_data_fetcher/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Expected Terminal Output
```text
--- Fetched To-Do Item ---
ID: 1
User ID: 1
Title: delectus aut autem
Completed: false