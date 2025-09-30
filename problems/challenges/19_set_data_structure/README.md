# Challenge #19: Set Data Structure

## Objective
Write a Go program to implement a simple Set data structure for strings. A Set is a collection that contains no duplicate elements. This challenge will test your understanding of maps, which are the perfect tool in Go to build a Set.

## Difficulty
Medium

## Concepts Tested
* Maps (`map[string]struct{}`)
* Structs (`struct`)
* Methods with Pointer Receivers
* Slices and Loops

## Rules/Logic
1.  Define a struct named `Set` with one field: `elements` (a map of type `map[string]struct{}`).
    * **Pro Tip**: We use an empty struct `struct{}` as the map's value because it takes up zero memory, making it the most efficient way to use a map for storing just keys.
2.  Implement the following three methods for the `Set` struct. Methods that modify the set must use a pointer receiver (`*Set`).
    * `Add(element string)`: Adds a new string to the set. If the element is already present, the set remains unchanged.
    * `Remove(element string)`: Removes a string from the set.
    * `Contains(element string) bool`: Returns `true` if the element is in the set, and `false` otherwise.
3.  Create a helper function `NewSet() *Set` that initializes and returns a new, empty set.
4.  In your `main` function:
    * Create a new set using your constructor.
    * Add several strings to it, including some duplicates, to show that only unique items are stored.
    * Check for the existence of an element using `Contains`.
    * Remove an element.
    * Try to check for the removed element again.
    * Print the final elements in the set (since maps are unordered, you can just loop through the map and print the keys).

## Your Tasks
1.  Create a new directory: `problems/challenges/19_set_data_structure/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
set := NewSet()

set.Add("apple")
set.Add("banana")
set.Add("apple") // Add a duplicate

fmt.Println("Contains 'apple'?", set.Contains("apple"))
fmt.Println("Contains 'grape'?", set.Contains("grape"))

set.Remove("apple")
fmt.Println("Contains 'apple' after removing?", set.Contains("apple"))

fmt.Println("Final elements in the set:")
// You can create a method on the Set to list elements if you like,
// or just loop through the map in main.
for element := range set.elements {
    fmt.Println("-", element)
}

// Expected Terminal Output (the order of final elements may vary):
Contains 'apple'? true
Contains 'grape'? false
Contains 'apple' after removing? false
Final elements in the set:
- banana