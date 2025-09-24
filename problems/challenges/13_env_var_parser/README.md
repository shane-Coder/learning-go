# Challenge #13: Environment Variable Parser

## Objective
Write a Go program that takes a slice of strings, where each string represents an environment variable (e.g., `"KEY=VALUE"`), and parses them into a map. The program should handle errors for malformed strings.

## Difficulty
Easy-Medium

## Concepts Tested
* Maps (`map[string]string`)
* String Manipulation (`strings.SplitN`)
* Error Handling (`errors.New`)
* Slices and Loops (`[]string`, `for...range`)

## Rules/Logic
1.  Create a function `parseEnvVars(envVars []string) (map[string]string, error)`.
2.  This function should take a slice of strings as input.
3.  Inside the function, create an empty map to store the parsed key-value pairs.
4.  Loop through each string in the input slice. For each string:
    * Use `strings.SplitN(s, "=", 2)` to split the string into a key and a value at the first `=` sign.
    * **Error Handling**: If the split does not result in exactly two parts, it's a malformed entry. The function should immediately return `nil` for the map and a new error (e.g., `errors.New("malformed entry: " + s)`).
    * If the split is successful, add the key and value to your map.
5.  If the loop completes without any errors, return the populated map and a `nil` error.
6.  In your `main` function, create a sample slice of environment variable strings, including at least one malformed entry to test your error handling.
7.  Call your `parseEnvVars` function, check for an error, and print either the resulting map or the error message.

## Your Tasks
1.  Create a new directory: `problems/challenges/13_env_var_parser/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
vars := []string{
    "PORT=8080",
    "DB_USER=root",
    "DB_PASS=password123",
    "INVALID_VAR", // A malformed entry
}

config, err := parseEnvVars(vars)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Parsed Config: %v\n", config)
}

// Also test a valid case
varsValid := []string{
    "HOST=localhost",
    "TIMEOUT=30s",
}

configValid, err := parseEnvVars(varsValid)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Parsed Config: %v\n", configValid)
}

// Expected Terminal Output:
Error: malformed entry: INVALID_VAR
Parsed Config: map[HOST:localhost TIMEOUT:30s]