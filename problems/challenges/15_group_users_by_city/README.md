# Challenge #15: Grouping Users by City

## Objective
Write a Go program that takes a slice of `User` structs and reorganizes them into a map where the keys are city names and the values are slices of users who live in that city.

## Difficulty
Medium

## Concepts Tested
* Maps with Slice Values (`map[string][]User`)
* Structs (`struct`)
* Slices and Loops (`[]User`, `for...range`)
* The `append` function

## Rules/Logic
1.  Define a struct named `User` with three fields: `ID` (int), `Name` (string), and `City` (string).
2.  Create a function `groupUsersByCity(users []User) map[string][]User`.
3.  Inside this function, create an empty map that will hold the grouped users. The map's type should be `map[string][]User`.
4.  Loop through the input slice of `User` structs. For each user:
    * Get the user's city.
    * Use the city as a key to access the map.
    * Append the current user to the slice of users for that city.
5.  Return the final map.
6.  In your `main` function:
    * Create a slice of `User` structs with users from a few different cities.
    * Call your `groupUsersByCity` function.
    * Loop through the resulting map and print the users for each city in a formatted way.

## Your Tasks
1.  Create a new directory: `problems/challenges/15_group_users_by_city/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
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

// Expected Terminal Output (the order of cities may vary):
City: Delhi
  - User: Shivam (ID: 1)
  - User: Bob (ID: 3)
City: Mumbai
  - User: Alice (ID: 2)
  - User: Diana (ID: 5)
City: Bangalore
  - User: Charlie (ID: 4)