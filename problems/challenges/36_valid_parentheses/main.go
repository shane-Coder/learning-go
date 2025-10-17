package main

import (
	"errors"
	"fmt"
)

// --- Simple Stack Implementation for Runes ---
type Stack struct {
	elements []rune
}

func (s *Stack) Push(element rune) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// --- End of Stack Implementation ---

// isValid checks if a string of parentheses is valid.
func isValid(s string) bool {
	stack := Stack{}

	// A map to quickly find the matching opening bracket for a closing one.
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Loop through each character (rune) in the input string.
	for _, char := range s {
		switch char {
		case '(', '{', '[': // If char is one of these opening brackets
			stack.Push(char)
		case ')', '}', ']': // If char is one of these closing brackets
			// ... handle closing brackets ...
			if stack.IsEmpty() {
				return false
			}
			topElement, _ := stack.Pop()
			if matchingBrackets[char] != topElement {
				return false
			}
			// default:
			// Optional: If you want to explicitly ignore other characters,
			// you could add a default case, but it's not needed here.
		}
		// If it's an opening bracket, push it onto the stack.
		// if char == '(' || char == '{' || char == '[' {
		// 	stack.Push(char)
		// } else if char == ')' || char == '}' || char == ']' { // If it's a closing bracket...
		// 	// Check if the stack is empty. If so, there's no matching open bracket.
		// 	if stack.IsEmpty() {
		// 		return false
		// 	}

		// 	// Pop the top element from the stack.
		// 	topElement, _ := stack.Pop() // We can ignore the error because we checked IsEmpty().

		// 	// Check if the popped element is the correct opening bracket.
		// 	if matchingBrackets[char] != topElement {
		// 		return false
		// 	}
		// }
		// Ignore any other characters that are not brackets.
	}

	// After looping through the entire string, the stack must be empty
	// for the string to be valid. If it's not empty, some brackets were left open.
	return stack.IsEmpty()
}

func main() {
	// Sample code in your main function:
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
	fmt.Println(isValid(""))       // true
	fmt.Println(isValid("["))      // false
	fmt.Println(isValid("]"))      // false
}
