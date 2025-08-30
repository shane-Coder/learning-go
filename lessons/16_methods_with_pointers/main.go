package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// This method uses a POINTER RECEIVER (*User).
// This gives the method the ability to modify the original User struct.
func (u *User) celebrateBirthday() {
	fmt.Printf("Celebrating %s's birthday...\n", u.Name)
	// Go automatically handles pointers here, so we can just write u.Age
	// instead of the more complex (*u).Age
	u.Age++
}

func main() {
	// Create an instance of User
	shivam := User{Name: "Shivam", Age: 25}
	fmt.Printf("Before birthday: %s is %d years old.\n", shivam.Name, shivam.Age)

	// Call the method that uses a pointer receiver
	// Notice we still call it the same way
	shivam.celebrateBirthday()

	// The original 'shivam' struct has been changed
	fmt.Printf("After birthday: %s is now %d years old.\n", shivam.Name, shivam.Age)
}
