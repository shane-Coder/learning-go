# Challenge #12: Basic Web Server with Multiple Routes

## Objective
Write a simple Go web server that responds with different messages on different URL paths (routes). This is the most fundamental building block of any web application or API.

## Difficulty
Easy

## Concepts Tested
* HTTP Server (`net/http`)
* Handling HTTP Requests (`http.ResponseWriter`, `*http.Request`)
* Routing (`http.HandleFunc`)
* Formatted Printing (`fmt.Fprintf`)

## Rules/Logic
1.  Create three simple handler functions:
    * `rootHandler`: This function will handle requests to the base URL (`/`). It should respond with the message "Welcome to the Home Page!".
    * `aboutHandler`: This function will handle requests to `/about`. It should respond with the message "This is the About Page.".
    * `contactHandler`: This function will handle requests to `/contact`. It should respond with the message "Contact us at support@example.com".
2.  Each handler function must have the signature `func(w http.ResponseWriter, r *http.Request)`.
3.  In your `main` function:
    * Register each of your handler functions to its corresponding URL path using `http.HandleFunc()`.
    * Print a message to the console indicating that the server is starting and which port it's listening on.
    * Start the server on port `8080` using `http.ListenAndServe()`.
    * Include error handling for `ListenAndServe` in case the port is already in use.

## Your Tasks
1.  Create a new directory: `problems/challenges/12_basic_web_server/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## How to Test
1.  Run your program with `go run main.go`.
2.  Open your web browser and visit the following URLs:
    * `http://localhost:8080/`
    * `http://localhost:8080/about`
    * `http://localhost:8080/contact`
    * `http://localhost:8080/other` (This should show a "404 page not found" error, which is correct).