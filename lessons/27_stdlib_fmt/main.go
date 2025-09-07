package main

import "fmt"

type User struct {
	Name string
	ID   int
}

func main() {
	user := User{Name: "Shivam", ID: 101}
	number := 42
	fruits := []string{"Apple", "Banana"}

	fmt.Println("--- Standard Verbs ---")
	fmt.Printf("Default value of user: %v\n", user)

	fmt.Println("\n--- Advanced Verbs ---")
	// %T prints the type of the variable
	fmt.Printf("Type of 'user': %T\n", user)
	fmt.Printf("Type of 'number': %T\n", number)

	// %+v prints the struct with field names
	fmt.Printf("User struct with field names: %+v\n", user)

	// %#v prints the Go-syntax representation
	fmt.Printf("Go-syntax for user: %#v\n", user)
	fmt.Printf("Go-syntax for fruits slice: %#v\n", fruits)
}
