# Challenge #26: Simple Banking System

## Objective
Write a Go program to simulate a simple banking system. You will create a `Bank` that can manage multiple `Account`s, allowing for deposits, withdrawals, and transfers between accounts, with robust error handling.

## Difficulty
Hard

## Concepts Tested
* Structs and Maps (`map[string]*Account`)
* Methods with Pointer Receivers
* Error Handling (Custom Errors)
* Encapsulation

## Rules/Logic
1.  Define a struct named `Account` with fields for `ID` (string) and `Balance` (float64).
2.  Define a struct named `Bank` with one field: `accounts` (a map where keys are account IDs and values are pointers to `Account` structs: `map[string]*Account`).
3.  Create a `NewBank()` constructor that returns an initialized `Bank`.
4.  Implement the following methods for the `Bank`:
    * `CreateAccount(accountID string) (*Account, error)`: Creates a new account with a zero balance. It should return an error if an account with that ID already exists.
    * `GetAccount(accountID string) (*Account, error)`: Finds and returns an account. It should return an error if the account does not exist.
    * `Deposit(accountID string, amount float64) error`: Adds money to an account. It should return an error if the account doesn't exist or if the deposit amount is negative or zero.
    * `Withdraw(accountID string, amount float64) error`: Subtracts money from an account. It should return an error if the account doesn't exist, if the amount is negative or zero, or if there are insufficient funds.
    * `Transfer(fromID, toID string, amount float64) error`: Moves money from one account to another. This method should handle all the same errors as `Withdraw` and `Deposit`. **Crucially, if the withdrawal from `fromID` fails, the deposit to `toID` should not happen.**
5.  In your `main` function, simulate a series of banking operations to test all your methods and their error conditions.

## Your Tasks
1.  Create a new directory: `problems/challenges/26_simple_banking_system/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
bank := NewBank()

bank.CreateAccount("shivam")
bank.CreateAccount("alice")

bank.Deposit("shivam", 100.0)
fmt.Println(bank.GetAccount("shivam")) // Should show balance 100

err := bank.Withdraw("shivam", 20.0) // shivam's balance is now 80
if err != nil { fmt.Println("Error:", err) }

err = bank.Transfer("shivam", "alice", 30.0) // shivam: 50, alice: 30
if err != nil { fmt.Println("Error:", err) }

fmt.Println(bank.GetAccount("shivam")) // Should show balance 50
fmt.Println(bank.GetAccount("alice"))  // Should show balance 30

err = bank.Withdraw("shivam", 100.0) // Test insufficient funds
if err != nil { fmt.Println("Error:", err) }

// Expected Terminal Output:
&{shivam 100}
&{shivam 50}
&{alice 30}
Error: insufficient funds for withdrawal