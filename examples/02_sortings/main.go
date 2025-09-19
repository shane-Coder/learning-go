package main

import (
	"fmt"
	"sort"
)

// Profile is the custom struct we want to sort.
type Profile struct {
	Name string
	Age  int
}

// --- Method 2: Implementing sort.Interface ---

// ByAge is a custom type for a slice of Profiles, which we will teach how to sort by age.
type ByAge []Profile

// Len is the number of elements in the collection.
func (a ByAge) Len() int {
	return len(a)
}

// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less reports whether the element with index i should be sorted before the element with index j.
func (a ByAge) Less(i, j int) bool {
	// This is our sorting logic: sort by Age in ascending order.
	return a[i].Age < a[j].Age
}

func main() {
	profiles := []Profile{
		{Name: "Shivam", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 22},
	}

	// --- Method 1: Using sort.Slice ---
	// We make a copy to show that sort.Slice sorts in-place.
	profilesByName := make([]Profile, len(profiles))
	copy(profilesByName, profiles)

	fmt.Println("--- Sorting with sort.Slice ---")
	// Here, we provide a custom "less" function to sort by Name alphabetically.
	sort.Slice(profilesByName, func(i, j int) bool {
		return profilesByName[i].Name < profilesByName[j].Name
	})
	fmt.Printf("Profiles sorted by name: %+v\n", profilesByName)

	// --- Method 2: Using sort.Interface ---
	fmt.Println("\n--- Sorting with sort.Interface (by Age) ---")
	// We call sort.Sort on our original slice after converting it to the ByAge type.
	sort.Sort(ByAge(profiles))
	fmt.Printf("Profiles sorted by age: %+v\n", profiles)
}
