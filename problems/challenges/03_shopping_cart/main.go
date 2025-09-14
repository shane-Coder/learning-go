package main

import (
	"fmt"
)

type Item struct {
	Name  string
	Price float64
}

type Cart struct {
	Items []Item
}

func (c *Cart) addItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c Cart) calculateTotal() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Price
	}
	return total
}

func (c Cart) displayCart() {
	if len(c.Items) == 0 {
		fmt.Println("Your shopping cart is empty.")
		return
	}
	fmt.Println("--- Shopping Cart ---")
	for _, item := range c.Items {
		fmt.Printf("Item: %s, Price: $%.2f\n", item.Name, item.Price)
	}
	fmt.Println("---------------------")
	total := c.calculateTotal()
	fmt.Printf("Total Price: $%.2f\n", total)
}

func NewCart() *Cart {
	return &Cart{Items: []Item{}}
}

func main() {
	cart := NewCart() // Use the constructor

	item1 := Item{Name: "Laptop", Price: 1200.00}
	item2 := Item{Name: "Mouse", Price: 25.50}
	item3 := Item{Name: "Keyboard", Price: 75.00}
	item4 := Item{Name: "Monitor", Price: 300.00}

	cart.addItem(item1)
	cart.addItem(item2)
	cart.addItem(item3)
	cart.addItem(item4)

	cart.displayCart()
}
