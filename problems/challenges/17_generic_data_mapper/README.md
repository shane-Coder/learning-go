# Challenge #17: Generic Data Mapper

## Objective
Write a Go program that uses a **generic function** to map a slice of one type to a slice of another type. This is a common operation in data processing pipelines, where you need to transform data from a source format to a destination format.

## Difficulty
Hard

## Concepts Tested
* Generics (with multiple type parameters)
* Functions as Arguments
* Slices
* Structs

## Rules/Logic
1.  Define a source struct named `Order` with fields for `ID` (int), `Amount` (float64), and `CustomerName` (string).
2.  Define a destination struct named `Receipt` with fields for `TransactionID` (int) and `AmountPaid` (float64).
3.  Create a **generic function** `MapSlice[T, U any](source []T, mapper func(T) U) []U`.
    * This function takes two type parameters: `T` for the source type and `U` for the destination type.
    * It takes two arguments: a slice of the source type (`source []T`) and a `mapper` function.
    * The `mapper` function is responsible for converting one element of type `T` into one element of type `U`.
4.  Inside your `MapSlice` function, loop through the `source` slice. For each element, call the `mapper` function and append the result to a new slice of type `[]U`. Return the new slice.
5.  In your `main` function:
    * Create a slice of `Order` structs.
    * Call your generic `MapSlice` function. For the `mapper` argument, provide a function that takes an `Order` and returns a `Receipt`.
    * Print the resulting slice of `Receipt` structs.

## Your Tasks
1.  Create a new directory: `problems/challenges/17_generic_data_mapper/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
orders := []Order{
    {ID: 1, Amount: 150.50, CustomerName: "Shivam"},
    {ID: 2, Amount: 75.25, CustomerName: "Alice"},
    {ID: 3, Amount: 200.00, CustomerName: "Bob"},
}

// The mapper function defines the conversion logic.
orderToReceiptMapper := func(order Order) Receipt {
    return Receipt{
        TransactionID: order.ID,
        AmountPaid:    order.Amount,
    }
}

receipts := MapSlice(orders, orderToReceiptMapper)

fmt.Printf("Generated Receipts: %+v\n", receipts)

// Expected Terminal Output:
Generated Receipts: [{TransactionID:1 AmountPaid:150.5} {TransactionID:2 AmountPaid:75.25} {TransactionID:3 AmountPaid:200}]