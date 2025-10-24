# Challenge #43: Queue Data Structure

## Objective
Write a Go program to implement a simple Queue data structure for integers. A queue follows the **FIFO (First-In, First-Out)** principle, meaning the first element added is the first one to be removed. Think of it like a line of people waiting.

## Difficulty
Medium

## Concepts Tested
* Structs (`struct`)
* Methods with Pointer Receivers
* Slices and Slice Manipulation
* Error Handling (`errors.New`)

## Rules/Logic
1.  Define a struct named `Queue` with one field: `elements` (a slice of integers `[]int`).
2.  Implement the following three methods for the `Queue` struct. Methods that modify the queue must use a pointer receiver (`*Queue`).
    * `Enqueue(element int)`: Adds a new element to the **back** of the queue.
    * `Dequeue() (int, error)`: Removes and returns the element from the **front** of the queue. If the queue is empty, it should return a zero value for the int and an error (e.g., `errors.New("queue is empty")`).
    * `Peek() (int, error)`: Returns the element at the **front** of the queue **without** removing it. It should also return an error if the queue is empty.
3.  Implement a helper method:
    * `IsEmpty() bool`: Returns `true` if the queue has no elements.
4.  In your `main` function, create a new queue and perform a sequence of `Enqueue` and `Dequeue` operations to demonstrate that it's working correctly (FIFO). Be sure to test the error case by trying to dequeue from an empty queue.

## Your Tasks
1.  Create a new directory: `problems/challenges/43_queue_data_structure/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
queue := Queue{}
fmt.Println("Is queue empty?", queue.IsEmpty()) // true

queue.Enqueue(10)
queue.Enqueue(20)
queue.Enqueue(30)

frontElement, _ := queue.Peek()
fmt.Println("Front element is:", frontElement) // 10

dequeuedElement, _ := queue.Dequeue()
fmt.Println("Dequeued:", dequeuedElement) // 10 (First In)

dequeuedElement, _ = queue.Dequeue()
fmt.Println("Dequeued:", dequeuedElement) // 20

fmt.Println("Is queue empty?", queue.IsEmpty()) // false

dequeuedElement, _ = queue.Dequeue()
fmt.Println("Dequeued:", dequeuedElement) // 30

_, err := queue.Dequeue() // Try to dequeue from an empty queue
if err != nil {
    fmt.Println("Error:", err)
}

// Expected Terminal Output:
Is queue empty? true
Front element is: 10
Dequeued: 10
Dequeued: 20
Is queue empty? false
Dequeued: 30
Error: queue is empty