package main

import "fmt"

func main() {
	// Create a map using make(). Keys are strings, values are ints.
	userAges := make(map[string]int)

	// Add key-value pairs
	userAges["Shivam"] = 25
	userAges["Alice"] = 30
	userAges["Bob"] = 35
	fmt.Println("Initial map:", userAges)

	// Access a value
	shivamAge := userAges["Shivam"]
	fmt.Printf("Shivam's age is %d\n", shivamAge)

	// The "comma ok" idiom to check for a key
	age, ok := userAges["Charlie"]
	if ok {
		fmt.Printf("Charlie's age is %d\n", age)
	} else {
		fmt.Println("Charlie's age not found.")
	}

	// Delete a key-value pair
	delete(userAges, "Bob")
	fmt.Println("Map after deleting Bob:", userAges)

	// You can also use for...range on a map!
	fmt.Println("--- All Users ---")
	for name, currentAge := range userAges {
		fmt.Printf("%s is %d years old.\n", name, currentAge)
	}
}
