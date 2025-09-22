# Challenge #11: Data Transformation Pipeline

## Objective
Write a Go program that processes a slice of `User` structs. You will filter out inactive users and then transform the remaining active users into a different struct format (`ActiveUserDTO`), simulating a data transformation pipeline.

## Difficulty
Medium

## Concepts Tested
* Structs (`struct`)
* Slices (`[]struct`)
* Loops (`for...range`)
* Functions as Transformation Steps
* Conditionals (`if`)

## Rules/Logic
1.  Define a struct named `User` with three fields: `ID` (int), `Name` (string), and `IsActive` (bool).
2.  Define a second struct named `ActiveUserDTO` (Data Transfer Object) with two fields: `ID` (int) and `DisplayName` (string).
3.  Create a function `filterInactiveUsers(users []User) []User`. This function should take a slice of `User` structs and return a new slice containing only the users where `IsActive` is `true`.
4.  Create a second function `transformUsersToDTOs(users []User) []ActiveUserDTO`. This function should take a slice of `User` structs and return a new slice of `ActiveUserDTO` structs. The `DisplayName` in the DTO should be the `Name` from the original `User` struct.
5.  In your `main` function:
    * Create an initial slice of `User` structs with a mix of active and inactive users.
    * Pass this slice through your `filterInactiveUsers` function to get a new slice of only active users.
    * Pass the result of the filtering step through your `transformUsersToDTOs` function to get the final slice of DTOs.
    * Print the final slice of `ActiveUserDTO`s.

## Your Tasks
1.  Create a new directory: `problems/challenges/11_data_transformation/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
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

// Expected Terminal Output:
Final DTOs: [{ID:1 DisplayName:Shivam} {ID:3 DisplayName:Bob} {ID:4 DisplayName:Charlie}]