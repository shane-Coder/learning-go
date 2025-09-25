package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]
	s.elements = s.elements[:topIndex] // Remove the top element
	return topElement, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func main() {
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
}
