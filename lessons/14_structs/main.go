package main

import "fmt"

// 1. Define the blueprint for a User
type User struct {
	Name     string
	Age      int
	IsActive bool
	Email    string
}

func main() {
	// 2. Create an instance of the User struct
	user1 := User{
		Name:     "Shivam Omer",
		Age:      25,
		IsActive: true,
		Email:    "shivam@example.com",
	}
	user2 := User{
		Name:     "Alice Johnson",
		Age:      30,
		IsActive: false,
		Email:    "alice@example.com",
	}

	// You can print the whole struct
	fmt.Println("User 1 details:", user1)
	fmt.Println("User 2 details:", user2)

	// 3. Access individual fields using dot notation
	fmt.Printf("%s is %d years old.\n", user1.Name, user1.Age)
	fmt.Printf("%s is %d years old.\n", user2.Name, user2.Age)

	// You can also modify a field's value
	user1.IsActive = false
	fmt.Println("User 1 after update:", user1)
	user2.Name = "Alice Smith"
	fmt.Println("User 2 after update:", user2)
}
