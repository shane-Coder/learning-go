package main

import (
	"errors"
	"fmt"
)

// Queue is a simple FIFO (First-In, First-Out) data structure.
type Queue struct {
	elements []int
}

// Enqueue adds an element to the back of the queue.
// Uses a pointer receiver (*Queue) because it modifies the slice.
func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element) // Append adds to the end (back).
}

// Dequeue removes and returns the element from the front of the queue.
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	// Get the element at the front (index 0).
	element := q.elements[0]
	// Reslice to remove the front element.
	q.elements = q.elements[1:]
	return element, nil
}

// Peek returns the element at the front of the queue without removing it.
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	// Return the element at the front (index 0).
	return q.elements[0], nil
}

// IsEmpty checks if the queue has any elements.
// This should also be a method on the Queue.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	// Initialize the queue correctly.
	// You don't *need* a NewQueue function, just creating the struct is fine.
	queue := Queue{} // The 'elements' slice starts as nil, which is okay.

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
}
