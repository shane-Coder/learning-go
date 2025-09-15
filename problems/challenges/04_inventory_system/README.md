# Challenge #4: Simple Inventory System

## Objective
Write a Go program to manage a simple product inventory. You will create different product types (e.g., a physical book and a digital download) that all satisfy a common `Product` interface, allowing them to be managed in a single inventory list.

## Difficulty
Medium

## Concepts Tested
* Interfaces (`interface`)
* Structs (`struct`)
* Methods
* Slices of Interfaces (`[]Product`)
* Type Assertions or Type Switches

## Rules/Logic
1.  Define an interface named `Product` with two methods:
    * `getName() string`: Returns the name of the product.
    * `getPrice() float64`: Returns the price of the product.
2.  Create two structs that will represent different kinds of products:
    * `Book`: Should have fields like `Name` (string), `Author` (string), and `Price` (float64).
    * `DigitalDownload`: Should have fields like `Name` (string), `FileSizeMB` (int), and `Price` (float64).
3.  Implement the `Product` interface for both the `Book` and `DigitalDownload` structs.
4.  Create a function `calculateInventoryValue(inventory []Product) float64`. This function should take a slice of `Product` interfaces, loop through them, and return the total value of all products in the inventory.
5.  In your `main` function:
    * Create a slice of `Product` interfaces.
    * Add at least one `Book` and one `DigitalDownload` to this slice.
    * Call `calculateInventoryValue` with your slice and print the total value.

## Your Tasks
1.  Create a new directory: `problems/challenges/04_inventory_system/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md` file.
4.  Write your solution in the `main.go` file.

## Sample Input/Output

```go
// Sample code in your main function:
inventory := []Product{
    Book{Name: "The Go Programming Language", Author: "Donovan & Kernighan", Price: 45.50},
    DigitalDownload{Name: "GoLand IDE License", FileSizeMB: 500, Price: 199.00},
    Book{Name: "Clean Architecture", Author: "Robert C. Martin", Price: 35.99},
}

totalValue := calculateInventoryValue(inventory)
fmt.Printf("Total inventory value: $%.2f\n", totalValue)

// Expected Terminal Output:
Total inventory value: $280.49
