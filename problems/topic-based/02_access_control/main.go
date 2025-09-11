/*
Problem Statement #2: User Role Access Control

Objective:
	Write a Go program that checks if a user has permission to access a protected resource based on their role, login status, and access level.

Rules/Logic:
You will have three variables:

	userRole (a string: e.g., "admin", "editor", "viewer")

	isLoggedIn (a bool: true or false)

	accessLevel (an int: e.g., 4, 5, 10)

Your program should print one of the following messages: "Access Granted", "Access Denied", or "Partial Access Granted".

	1. Use a switch statement on the userRole:

		case "admin": They always get full access. Print "Access Granted".

		case "editor":

			if their accessLevel is 5 or higher, they get full access. Print "Access Granted"

			else if they are just isLoggedIn (but their level is too low), they get partial access. Print "Partial Access Granted"

			otherwise, print "Access Denied"

		case "viewer": They only need to be logged in (isLoggedIn must be true). Print "Access Granted" or "Access Denied".

		default: Any other role has no access. Print "Access Denied".

--- Terminal Output ---

	// Test Case 1: Editor with low access level but logged in
	UserRole: editor, IsLoggedIn: true, AccessLevel: 3
	Partial Access Granted

	// Test Case 2: Editor with high access level
	UserRole: editor, IsLoggedIn: false, AccessLevel: 10
	Access Granted

	// Test Case 3: Admin always has access
	UserRole: admin, IsLoggedIn: false, AccessLevel: 0
	Access Granted

	// Test Case 4: Viewer needs to be logged in
	UserRole: viewer, IsLoggedIn: true, AccessLevel: 1
	Access Granted

	// Test Case 5: Default role is denied
	UserRole: sender, IsLoggedIn: true, AccessLevel: 10
	Access Denied
*/

package main

import "fmt"

func main() {
	fmt.Println("Welcome to the User Role Access Control System!")

	userRole := "sender"
	isLoggedIn := true
	accessLevel := 10

	fmt.Println("UserRole:", userRole, ", IsLoggedIn:", isLoggedIn, ", AccessLevel:", accessLevel)

	// code logic
	switch userRole {
	case "admin":
		print("Access Granted")
	case "editor":
		if accessLevel >= 5 {
			print("Access Granted")
		} else if isLoggedIn {
			print("Partial Access Granted")
		} else {
			print("Access Denied")
		}
	case "viewer":
		if isLoggedIn {
			print("Access Granted")
		} else {
			print("Access Denied")
		}
	default:
		print("Access Denied")
	}
	fmt.Println() // for new line
}
