package main

import "fmt"

type Order struct {
	ID           int
	Amount       float64
	CustomerName string
}

type Receipt struct {
	TransactionID int
	AmountPaid    float64
}

func MapSlice[T, U any](source []T, mapper func(T) U) []U {
	// var result []U // Initialize an empty slice of the destination type U.
	result := make([]U, 0, len(source)) // A slightly more performant way to initialize

	for _, value := range source {
		// For each element, call the mapper function to transform it, and then append the result to our new slice.
		result = append(result, mapper(value))
	}
	return result
}

func main() {
	// --- Test Case 1: The Original Challenge (Order -> Receipt) ---
	fmt.Println("--- Test Case 1: Mapping Orders to Receipts ---")
	orders := []Order{
		{ID: 1, Amount: 150.50, CustomerName: "Shivam"},
		{ID: 2, Amount: 75.25, CustomerName: "Alice"},
		{ID: 3, Amount: 200.00, CustomerName: "Bob"},
	}

	// The mapper function defines the conversion logic from Order to Receipt.
	orderToReceiptMapper := func(order Order) Receipt {
		return Receipt{
			TransactionID: order.ID,
			AmountPaid:    order.Amount,
		}
	}

	receipts := MapSlice(orders, orderToReceiptMapper)
	fmt.Printf("Generated Receipts: %+v\n", receipts)

	// --- Test Case 2: An Empty Slice ---
	fmt.Println("\n--- Test Case 2: Handling an Empty Slice ---")
	emptyOrders := []Order{}
	emptyReceipts := MapSlice(emptyOrders, orderToReceiptMapper)
	fmt.Printf("Result from empty slice (should be empty): %+v\n", emptyReceipts)

	// --- Test Case 3: A Different Transformation (Order -> string) ---
	fmt.Println("\n--- Test Case 3: Mapping Orders to Strings ---")
	// This mapper transforms an Order into a summary string.
	orderToSummaryMapper := func(order Order) string {
		return fmt.Sprintf("Order %d for %s was $%.2f", order.ID, order.CustomerName, order.Amount)
	}

	summaries := MapSlice(orders, orderToSummaryMapper)
	fmt.Println("Generated Summaries:")
	for _, s := range summaries {
		fmt.Printf("  - %s\n", s)
	}
}
