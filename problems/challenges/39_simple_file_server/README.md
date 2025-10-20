# Challenge #39: Simple File Server

## Objective
Write a Go program that acts as a simple web server to serve static files (like HTML, CSS, or images) from a specific directory.

## Difficulty
Medium

## Concepts Tested
* HTTP Server (`net/http`)
* Serving Static Files (`http.FileServer`, `http.Dir`)
* Basic Routing

## Rules/Logic
1.  Create a directory named `public` in your project folder. Inside this `public` directory, create a simple HTML file named `index.html`. You can put basic content like `<h1>Hello from Go File Server!</h1>` in it.
2.  In your `main.go` file, create a file server handler using `http.FileServer(http.Dir("./public"))`.
    * `http.Dir("./public")` tells Go to look for files relative to the `public` directory.
    * `http.FileServer()` creates a handler that knows how to serve those files.
3.  Register this file server handler to handle requests for the root path (`/`) using `http.Handle()`.
    * **Note**: We use `http.Handle()` here, not `http.HandleFunc()`, because `http.FileServer` returns a handler object directly.
4.  In your `main` function:
    * Print a message indicating the server is starting and the port.
    * Start the server on port `8080` using `http.ListenAndServe()`.
    * Include error handling for `ListenAndServe`.

## Your Tasks
1.  Create a new directory: `problems/challenges/39_simple_file_server/`.
2.  Inside this directory, create `README.md`, `main.go`, and a subdirectory named `public`.
3.  Inside the `public` subdirectory, create `index.html` with some basic HTML content.
4.  Copy and paste this problem statement into your `README.md`.
5.  Write your solution in `main.go`.

## How to Test
1.  Run your program with `go run main.go`.
2.  Open your web browser and visit `http://localhost:8080/`.
3.  You should see the content of your `index.html` file displayed in the browser.