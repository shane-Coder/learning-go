package main

import "fmt"

type User struct {
	ID   int
	Name string
	City string
}

func groupUsersByCity(users []User) map[string][]User {
	var result = make(map[string][]User)
	for _, user := range users {
		city := user.City
		result[city] = append(result[city], user)
	}
	return result
}

func main() {
	users := []User{
		{ID: 1, Name: "Shivam", City: "Delhi"},
		{ID: 2, Name: "Alice", City: "Mumbai"},
		{ID: 3, Name: "Bob", City: "Delhi"},
		{ID: 4, Name: "Charlie", City: "Bangalore"},
		{ID: 5, Name: "Diana", City: "Mumbai"},
	}

	groupedUsers := groupUsersByCity(users)

	for city, cityUsers := range groupedUsers {
		fmt.Printf("City: %s\n", city)
		for _, user := range cityUsers {
			fmt.Printf("  - User: %s (ID: %d)\n", user.Name, user.ID)
		}
	}
}
