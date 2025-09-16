# Challenge #5: User Profile Manager with Custom Errors

## Objective
Write a Go program that manages user profiles and returns specific, custom errors for different invalid input scenarios. This is a common requirement for any application that handles user data.

## Difficulty
Medium

## Concepts Tested
* Structs (`struct`)
* Custom Error Types
* Error Handling (`if err != nil`)
* Functions
* The `errors` Package

## Rules/Logic
1.  Define a struct named `Profile` with two fields: `Username` (string) and `Age` (int).
2.  Define two custom error types:
    * `ValidationError`: A struct with a `Message` field (string).
    * `DuplicateUserError`: A struct with a `Username` field (string).
3.  Both custom error structs must implement the built-in `error` interface by having an `Error() string` method.
4.  Create a "database" using a simple map to store user profiles: `database := make(map[string]Profile)`.
5.  Create a function `createProfile(username string, age int) (*Profile, error)`. This function will handle the validation and creation logic.
    * **Validation**:
        * If the `username` is less than 3 characters long, return a `ValidationError`.
        * If the `age` is less than 18, return a `ValidationError`.
    * **Duplicate Check**:
        * If the `username` already exists in your map "database," return a `DuplicateUserError`.
    * **Success**:
        * If all checks pass, create a new `Profile`, add it to the map, and return a pointer to the new profile and a `nil` error.
6.  In your `main` function, try to create several profiles to test all scenarios: a successful creation, a validation error (e.g., short username), and a duplicate user error.

## Your Tasks
1.  Create a new directory: `problems/challenges/05_user_profile_manager/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md` file.
4.  Write your solution in `main.go`. In your `main` function, use the `if err != nil` pattern to check for errors and print informative messages for each case.

## Sample Input/Output

```go
// Sample code in your main function to test the logic:
_, err := createProfile("shivam", 25)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Println("Profile for 'shivam' created successfully.")
}

_, err = createProfile("shivam", 30) // Test duplicate
if err != nil {
    fmt.Printf("Error: %v\n", err)
}

_, err = createProfile("ab", 22) // Test validation
if err != nil {
    fmt.Printf("Error: %v\n", err)
}

_, err = createProfile("alice", 17) // Test age validation
if err != nil {
    fmt.Printf("Error: %v\n", err)
}

// Output
cmd to run go run problems/challenges/05_user_profile_manager/main.go 
Profile for 'shivam' created successfully.
Error: Duplicate User Error: username 'shivam' already exists
Error: Validation Error: username must be at least 3 characters long
Error: Validation Error: age must be at least 18
saro@pavilion:~/LANKA/learning-go$ 