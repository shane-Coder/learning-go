# Challenge #14: Stack Data Structure

## Objective
Write a Go program to implement a simple Stack data structure for integers. A stack follows the **LIFO (Last-In, First-Out)** principle, meaning the last element added is the first one to be removed. Think of it like a stack of plates.

## Difficulty
Medium

## Concepts Tested
* Structs (`struct`)
* Methods with Pointer Receivers
* Slices and Slice Manipulation
* Error Handling (`errors.New`)

## Rules/Logic
1.  Define a struct named `Stack` with one field: `elements` (a slice of integers `[]int`).
2.  Implement the following three methods for the `Stack` struct. All methods that modify the stack must use a pointer receiver (`*Stack`).
    * `Push(element int)`: This method should add a new element to the top of the stack.
    * `Pop() (int, error)`: This method should remove and return the top element of the stack. If the stack is empty, it should return a zero value for the int and an error (e.g., `errors.New("stack is empty")`).
    * `Peek() (int, error)`: This method should return the top element of the stack **without** removing it. It should also return an error if the stack is empty.
3.  It's also good practice to add a helper method:
    * `IsEmpty() bool`: Returns `true` if the stack has no elements.
4.  In your `main` function, create a new stack and perform a sequence of `Push` and `Pop` operations to demonstrate that it's working correctly. Be sure to test the error case by trying to pop from an empty stack.

## Your Tasks
1.  Create a new directory: `problems/challenges/14_stack_data_structure/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
stack := Stack{}
fmt.Println("Is stack empty?", stack.IsEmpty()) // true

stack.Push(10)
stack.Push(20)
stack.Push(30)

topElement, _ := stack.Peek()
fmt.Println("Top element is:", topElement) // 30

poppedElement, _ := stack.Pop()
fmt.Println("Popped:", poppedElement) // 30

poppedElement, _ = stack.Pop()
fmt.Println("Popped:", poppedElement) // 20

fmt.Println("Is stack empty?", stack.IsEmpty()) // false

poppedElement, _ = stack.Pop()
fmt.Println("Popped:", poppedElement) // 10

_, err := stack.Pop() // Try to pop from an empty stack
if err != nil {
    fmt.Println("Error:", err)
}

// Expected Terminal Output:
Is stack empty? true
Top element is: 30
Popped: 30
Popped: 20
Is stack empty? false
Popped: 10
Error: stack is empty