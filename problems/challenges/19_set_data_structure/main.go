package main

import "fmt"

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{elements: make(map[string]struct{})}
}

func (s *Set) Add(element string) {
	s.elements[element] = struct{}{}
}

func (s *Set) Remove(element string) {
	delete(s.elements, element)
}

func (s *Set) Contains(element string) bool {
	_, ok := s.elements[element]
	return ok
}

func main() {
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
}
