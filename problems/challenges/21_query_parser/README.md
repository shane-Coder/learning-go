# Challenge #21: URL Query Parameter Parser

## Objective
Write a Go program with a function that takes a URL query string (the part of a URL after the `?`) and parses it into a map of keys and values. For example, the string `"key1=value1&key2=value2"` should become `map[string]string{"key1": "value1", "key2": "value2"}`.

## Difficulty
Medium

## Concepts Tested
* String Manipulation (`strings.Split`)
* Maps (`map[string]string`)
* Slices and Loops
* Error Handling (for malformed parameters)

## Rules/Logic
1.  Create a function `parseQuery(queryString string) (map[string]string, error)`.
2.  Inside the function, first split the `queryString` into pairs by the `&` character.
3.  Create an empty map to store the results.
4.  Loop through each pair string. For each pair:
    * Split the pair into a key and a value by the `=` character.
    * **Error Handling**: A valid pair must split into exactly two parts. If it doesn't, the function should immediately return `nil` and an error indicating which pair was malformed.
    * If the split is successful, add the key and value to your map.
5.  If the loop completes without any errors, return the populated map and a `nil` error.
6.  In your `main` function, test your `parseQuery` function with both a valid query string and one that contains a malformed parameter.

## Your Tasks
1.  Create a new directory: `problems/challenges/21_query_parser/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
query1 := "name=shivam&city=gurgaon&active=true"
params1, err1 := parseQuery(query1)
if err1 != nil {
    fmt.Println("Error:", err1)
} else {
    fmt.Println("Parsed Query 1:", params1)
}

query2 := "id=123&user=alice&malformed_param"
params2, err2 := parseQuery(query2)
if err2 != nil {
    fmt.Println("Error:", err2)
} else {
    fmt.Println("Parsed Query 2:", params2)
}


// Expected Terminal Output:
Parsed Query 1: map[active:true city:gurgaon name:shivam]
Error: malformed query parameter: malformed_param