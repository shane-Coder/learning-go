# Challenge #46: Flatten Nested Slice

## Objective
Write a Go program with a function that takes a potentially nested slice of integers (represented using `interface{}`) and flattens it into a single slice of integers.

## Difficulty
Medium-Hard

## Concepts Tested
* Interfaces (`interface{}`)
* Type Assertions or Type Switches
* Recursion (or Iteration with a Stack/Queue)
* Slices and `append`

## Rules/Logic
1.  Create a function `flatten(nested interface{}) []int`.
2.  The input `nested` will be an `interface{}` that holds either an integer (`int`) or a slice of `interface{}` (`[]interface{}`). This slice can contain more integers or more nested slices.
3.  The function must return a **new, flat slice** containing all the integers from the nested structure, in the order they appear.
4.  **Hint (Recursive Approach)**:
    * Use a type switch (or type assertion) on the input `nested`.
    * **Base Case**: If the input is an `int`, return a new slice containing just that integer.
    * **Recursive Step**: If the input is a slice (`[]interface{}`), iterate through its elements. For each element, recursively call `flatten` and append the results to a final result slice.
5.  In your `main` function, create a sample nested slice and test your `flatten` function.

## Your Tasks
1.  Create a new directory: `problems/challenges/46_flatten_nested_slice/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
nestedList := []interface{}{
    1,
    []interface{}{2, 3},
    4,
    []interface{}{
        5,
        []interface{}{6, 7},
    },
    8,
}

flattenedList := flatten(nestedList)
fmt.Println("Flattened list:", flattenedList)

// Test with a simpler case
simpleList := []interface{}{10, 20, 30}
flattenedSimple := flatten(simpleList)
fmt.Println("Flattened simple list:", flattenedSimple)

// Expected Terminal Output:
Flattened list: [1 2 3 4 5 6 7 8]
Flattened simple list: [10 20 30]