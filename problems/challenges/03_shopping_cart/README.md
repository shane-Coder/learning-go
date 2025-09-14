# Challenge #3: E-commerce Shopping Cart

## Objective
Write a Go program that simulates a simple e-commerce shopping cart. You will create structs for `Item` and `Cart` and then implement methods to add items to the cart and calculate the total price.

## Difficulty
Medium

## Concepts Tested
* Structs (`struct`)
* Methods (with pointer receivers)
* Slices of Structs (`[]Item`)
* Loops (`for`)
* Formatted Printing (`fmt.Printf`)

## Rules/Logic
1.  Define a struct named `Item` with two fields: `Name` (string) and `Price` (float64).
2.  Define a struct named `Cart` with one field: `Items` (a slice of `Item` structs).
3.  Create a method with a pointer receiver for the `Cart` struct named `addItem(item Item)`. This method should append the new item to the cart's `Items` slice.
4.  Create another method for the `Cart` struct named `calculateTotal() float64`. This method should loop through all the items in the cart and return their total price.
5.  Create a final method for the `Cart` named `displayCart()`. This method should print out each item's name and price, followed by the total price of the cart.

## Your Tasks
1.  Create a new directory: `problems/challenges/03_shopping_cart/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md` file.
4.  Write your solution in the `main.go` file.
5.  In your `main` function, create a new cart, add several different items to it, and then call the `displayCart` method to show the final result.

## Sample Input/Output

```go
// Sample code in your main function:
cart := Cart{}

item1 := Item{Name: "Laptop", Price: 1200.00}
item2 := Item{Name: "Mouse", Price: 25.50}
item3 := Item{Name: "Keyboard", Price: 75.00}

cart.addItem(item1)
cart.addItem(item2)
cart.addItem(item3)

cart.displayCart()

// Expected Terminal Output:
--- Shopping Cart ---
Item: Laptop, Price: $1200.00
Item: Mouse, Price: $25.50
Item: Keyboard, Price: $75.00
---------------------
Total Price: $1300.50
