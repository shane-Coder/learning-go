# Challenge #9: JSON Data Parser

## Objective
Write a Go program that parses a JSON string representing a user and their social media profiles, and then prints a formatted summary. This tests your ability to handle nested JSON objects and slices.

## Difficulty
Medium

## Concepts Tested
* JSON Unmarshalling (`encoding/json`)
* Structs (including nested structs)
* Slices of Structs
* Formatted Printing (`fmt.Printf`)

## Rules/Logic
1.  You will be given a JSON string that represents a user object. This object will contain basic fields and a nested slice of social media profiles.
2.  Define two structs to match the JSON structure:
    * `SocialProfile`: Should have fields for `Platform` (string) and `Username` (string).
    * `User`: Should have fields for `ID` (int), `Name` (string), `Email` (string), and `Profiles` (a slice of `SocialProfile` structs).
3.  Create a function `parseUserJSON(jsonString string) (User, error)` that:
    * Takes the JSON string as input.
    * Uses `json.Unmarshal` to parse the string into your `User` struct.
    * Returns the populated `User` struct and any error that occurred.
4.  In your `main` function:
    * Define the sample JSON string.
    * Call your `parseUserJSON` function and handle any potential errors.
    * If parsing is successful, print a formatted summary of the user's details, including a list of their social media profiles.

## Your Tasks
1.  Create a new directory: `problems/challenges/09_json_parser/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample JSON Input
```go
// This is the JSON string you'll parse in your main function.
const userJSON = `{
    "id": 101,
    "name": "Shivam",
    "email": "shivam@example.com",
    "profiles": [
        {
            "platform": "GitHub",
            "username": "shane-Coder"
        },
        {
            "platform": "LinkedIn",
            "username": "programmer-shivam"
        }
    ]
}`

--- User Profile ---
ID: 101
Name: Shivam
Email: shivam@example.com
Social Profiles:
  - Platform: GitHub, Username: shane-Coder
  - Platform: LinkedIn, Username: programmer-shivam