/*
Problem Statement #7: Bank Account Manager

Objective:
	Create a Go program that simulates basic bank account operations like deposits and withdrawals using a struct and methods.

Rules/Logic:

	1. Define a struct named Account with two fields: Owner (string) and Balance (float64).

	2. Create two methods with pointer receivers for the Account struct:

		(i) deposit(amount float64): This method should add the amount to the account's Balance.

		(ii) withdraw(amount float64): This method should subtract the amount from the Balance. Add a check to ensure the balance doesn't go below zero; if the withdrawal amount is too large, print an "Insufficient funds" message and don't change the balance.

	3. In your main function, create a new Account instance.

	4. Call the deposit and withdraw methods to show that they correctly modify the original account's balance. Print the balance after each operation.

---Terminal Output---
	// Test Case 1: Deposit and Withdraw
	Depositing ₹11111.10 into Ram's account.
	New Balance: ₹11111.10
	Withdrawing ₹10000.00 from Ram's account.
	New Balance: ₹1111.10
	Withdrawing ₹2000.00 from Ram's account.
	Withdrawal failed: Insufficient funds.

	// Test Case 2: Deposit and Withdraw
	Depositing ₹11111.10 into Shyam's account.
	New Balance: ₹11111.10
	Withdrawing ₹10000.00 from Shyam's account.
	New Balance: ₹1111.10
	Withdrawing ₹2000.00 from Shyam's account.
	Withdrawal failed: Insufficient funds.

*/

package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

// func (acc *Account) deposit(credit float64) float64 {
// 	fmt.Printf("Current balance in account owner %s is %.2f.\n", acc.Owner, acc.Balance)
// 	acc.Balance += credit
// 	fmt.Printf("After adding amount %.2f in balance in account owner %s is %.2f.\n", credit, acc.Owner, acc.Balance)
// 	return acc.Balance
// }

func (acc *Account) deposit(depositAmount float64) {
	fmt.Printf("Depositing ₹%.2f into %s's account.\n", depositAmount, acc.Owner)
	acc.Balance += depositAmount
	fmt.Printf("New Balance: ₹%.2f\n", acc.Balance)
}

// func (acc *Account) withdraw(debit float64) float64 {
// 	fmt.Printf("Current balance in account %s is %.2f.\n", acc.Owner, acc.Balance)
// 	if acc.Balance-debit < 0 {
// 		fmt.Printf("Insufficient funds in account %s.\n", acc.Owner)
// 	} else {
// 		acc.Balance -= debit
// 		fmt.Printf("After withdraw amount %.2f from balance in account %s is %.2f.\n", debit, acc.Owner, acc.Balance)
// 	}
// 	return acc.Balance
// }

func (acc *Account) withdraw(withdrawAmount float64) {
	fmt.Printf("Withdrawing ₹%.2f from %s's account.\n", withdrawAmount, acc.Owner)
	if withdrawAmount > acc.Balance {
		fmt.Println("Withdrawal failed: Insufficient funds.")
	} else {
		acc.Balance -= withdrawAmount
		fmt.Printf("New Balance: ₹%.2f\n", acc.Balance)
	}
}

func main() {
	account1 := Account{
		Owner:   "Ram",
		Balance: 0.00,
	}
	account2 := Account{
		Owner:   "Shyam",
		Balance: 21.75,
	}
	depositAmount := 11111.10
	withdrawAmount := 11111.11

	// fmt.Printf("Current balance in account %s is %.2f.\n", account1.Owner, account1.Balance)
	account1.deposit(depositAmount)
	// fmt.Printf("After adding amount %.2f from balance in account %s is %.2f.\n", credit, account1.Owner, account1.Balance)

	// fmt.Printf("Current balance in account %s is %.2f.\n", account2.Owner, account2.Balance)
	account2.deposit(depositAmount)
	// fmt.Printf("After adding amount %.2f from balance in account %s is %.2f.\n", credit, account2.Owner, account2.Balance)

	account1.withdraw(withdrawAmount)
	account2.withdraw(withdrawAmount)
}
