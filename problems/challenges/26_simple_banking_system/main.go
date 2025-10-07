package main

import (
	"errors"
	"fmt"
)

// Account struct holds the data for a single bank account.
type Account struct {
	ID      string
	Balance float64
}

// Bank struct holds all the accounts in a map.
type Bank struct {
	accounts map[string]*Account
}

// NewBank is a constructor that returns a new, initialized Bank.
func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

// CreateAccount creates a new account and adds it to the bank.
func (b *Bank) CreateAccount(accountID string) (*Account, error) {
	// Check if the account already exists.
	if _, ok := b.accounts[accountID]; ok {
		return nil, fmt.Errorf("account with ID '%s' already exists", accountID)
	}
	// Create the new account.
	newAccount := &Account{ID: accountID, Balance: 0.0}
	// Add it to the map.
	b.accounts[accountID] = newAccount
	return newAccount, nil
}

// GetAccount retrieves an account from the bank.
func (b *Bank) GetAccount(accountID string) (*Account, error) {
	account, ok := b.accounts[accountID]
	if !ok {
		return nil, fmt.Errorf("account with ID '%s' not found", accountID)
	}
	return account, nil
}

// Deposit adds money to an account.
func (b *Bank) Deposit(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	account, err := b.GetAccount(accountID)
	if err != nil {
		return err // Return the "not found" error from GetAccount.
	}
	account.Balance += amount
	return nil // Return nil for success.
}

// Withdraw removes money from an account.
func (b *Bank) Withdraw(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	account, err := b.GetAccount(accountID)
	if err != nil {
		return err
	}
	if account.Balance < amount {
		return errors.New("insufficient funds for withdrawal")
	}
	account.Balance -= amount
	return nil
}

// Transfer moves money between two accounts.
func (b *Bank) Transfer(fromID, toID string, amount float64) error {
	// First, try to withdraw the money from the source account.
	// If this fails (e.g., insufficient funds), the function will return the error
	// and the deposit will never happen. This makes the transaction safe.
	if err := b.Withdraw(fromID, amount); err != nil {
		return err
	}

	// If the withdrawal was successful, deposit the money into the destination account.
	if err := b.Deposit(toID, amount); err != nil {
		// In a real-world scenario, you would need to handle the case where the
		// deposit fails by rolling back the withdrawal. For this challenge, we assume it will succeed.
		return err
	}
	return nil
}

func main() {
	bank := NewBank()

	bank.CreateAccount("shivam")
	bank.CreateAccount("alice")

	fmt.Println("--- Initial State ---")
	shivamAcc, _ := bank.GetAccount("shivam")
	fmt.Printf("Shivam's account: %+v\n", shivamAcc)

	fmt.Println("\n--- After Deposit ---")
	bank.Deposit("shivam", 100.0)
	fmt.Printf("Shivam's account: %+v\n", shivamAcc)

	fmt.Println("\n--- After Withdrawal ---")
	err := bank.Withdraw("shivam", 20.0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Shivam's account: %+v\n", shivamAcc)

	fmt.Println("\n--- After Transfer ---")
	err = bank.Transfer("shivam", "alice", 30.0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	aliceAcc, _ := bank.GetAccount("alice")
	fmt.Printf("Shivam's account: %+v\n", shivamAcc)
	fmt.Printf("Alice's account: %+v\n", aliceAcc)

	fmt.Println("\n--- Testing Error Case ---")
	err = bank.Withdraw("shivam", 100.0) // Test insufficient funds
	if err != nil {
		fmt.Println("Error:", err)
	}
}
