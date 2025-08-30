package main

import "fmt"

// Our User struct blueprint
type User struct {
	Name string
	Age  int
}

// This is a METHOD on the User struct.
// The '(u User)' part is the "receiver".
// It means this function is now attached to the User type.
func (u User) greet() {
	fmt.Printf("Hello! My name is %s and I am %d years old.\n", u.Name, u.Age)
}

func main() {
	// Create an instance of User
	user1 := User{Name: "Shivam", Age: 25}

	// Call the method on the user1 instance using dot notation
	user1.greet()

	user2 := User{Name: "Alice", Age: 30}
	user2.greet()

	user3 := User{Name: "Bob", Age: 22}
	user3.greet()
}
