package main

import "fmt"

type Product interface {
	getName() string
	getPrice() float64
}

type Book struct {
	Name   string
	Author string
	Price  float64
}

type DigitalDownload struct {
	Name       string
	FileSizeMB int
	Price      float64
}

func (b Book) getName() string {
	return b.Name
}

func (b Book) getPrice() float64 {
	return b.Price
}

func (dd DigitalDownload) getName() string {
	return dd.Name
}

func (dd DigitalDownload) getPrice() float64 {
	return dd.Price
}

func calculateInventoryValue(inventory []Product) float64 {
	total := 0.0
	for _, price := range inventory {
		total += float64(price.getPrice())
	}
	return total
}

// printInventoryReport demonstrates a powerful feature of interfaces in Go: the type switch.
// This function iterates over a slice of generic Product interfaces and "unwraps"
// each one to discover its specific, concrete type (e.g., Book or DigitalDownload).
// This allows us to access type-specific fields (like Author or FileSizeMB) that are
// not part of the common interface, enabling detailed and type-specific logic.
func printInventoryReport(inventory []Product) {
	fmt.Println("--- Inventory Report ---")
	for _, p := range inventory {
		fmt.Printf("Name: %s, Price: $%.2f\n", p.getName(), p.getPrice())

		// This is a type switch
		switch product := p.(type) {
		case Book:
			fmt.Printf("  - Type: Book, Author: %s\n", product.Author)
		case DigitalDownload:
			fmt.Printf("  - Type: Digital Download, File Size: %d MB\n", product.FileSizeMB)
		}
	}
	fmt.Println("----------------------")
}

func main() {
	inventory := []Product{
		Book{Name: "The Go Programming Language", Author: "Donovan & Kernighan", Price: 45.50},
		DigitalDownload{Name: "GoLand IDE License", FileSizeMB: 500, Price: 199.00},
		Book{Name: "Clean Architecture", Author: "Robert C. Martin", Price: 35.99},
		DigitalDownload{Name: "E-book: Learning Go", FileSizeMB: 5, Price: 0.00},
	}

	// --- Call the new function ---
	printInventoryReport(inventory)

	totalValue := calculateInventoryValue(inventory)
	fmt.Printf("Total inventory value: $%.2f\n", totalValue)
}
