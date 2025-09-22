package main

import "fmt"

type User struct {
	ID       int
	Name     string
	IsActive bool
}

type ActiveUserDTO struct {
	ID          int
	DisplayName string
}

func filterInactiveUsers(users []User) []User {
	var activeUsers []User
	for _, user := range users {
		if user.IsActive {
			activeUsers = append(activeUsers, user)
		}
	}
	return activeUsers
}

func transformUsersToDTOs(users []User) []ActiveUserDTO {
	var displayUsers []ActiveUserDTO
	for _, user := range users {
		// displayUsers = append(displayUsers, ActiveUserDTO{
		// 	ID:          user.ID,
		// 	DisplayName: user.Name,
		// })
		dto := ActiveUserDTO{
			ID:          user.ID,
			DisplayName: user.Name,
		}
		displayUsers = append(displayUsers, dto)
	}
	return displayUsers
}

func main() {
	users := []User{
		{ID: 1, Name: "Shivam", IsActive: true},
		{ID: 2, Name: "Alice", IsActive: false},
		{ID: 3, Name: "Bob", IsActive: true},
		{ID: 4, Name: "Charlie", IsActive: true},
		{ID: 5, Name: "David", IsActive: false},
	}

	// 1. First pipeline step: Filtering
	activeUsers := filterInactiveUsers(users)

	// 2. Second pipeline step: Transformation
	activeUserDTOs := transformUsersToDTOs(activeUsers)

	fmt.Printf("Final DTOs: %+v\n", activeUserDTOs)
}
