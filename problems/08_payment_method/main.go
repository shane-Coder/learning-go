/*
Problem Statement #8: Payment Method Calculator

Objective:
	Write a Go program that can calculate the processing fee for different types of payment methods (like Credit Card and PayPal), each of which has a different fee calculation rule.

Rules/Logic:

	1. Define an interface named PaymentMethod that requires one method: calculateFee(amount float64) float64.

	2. Define two structs: CreditCard and PayPal.

		1. CreditCard should have a field for transactionFeePercentage (float64).

		2. PayPal should have a field for processingFee (a flat float64 amount).

	3. Implement the PaymentMethod interface for both structs:

		1. The CreditCard's calculateFee method should return the amount multiplied by its transactionFeePercentage (e.g., amount * 0.03 for 3%).

		2. The PayPal's calculateFee method should just return its flat processingFee, regardless of the transaction amount.

	4. Create a generic function getProcessingFee(p PaymentMethod, amount float64) that takes any type satisfying the PaymentMethod interface and prints the calculated fee.

	5. In your main function, create instances of CreditCard and PayPal, and pass both of them to your generic getProcessingFee function to show that it works with both types.

---Terminal Output---
*/

package main

import "fmt"

type PaymentMethod interface {
	calculateFee(amount float64) float64
	name() string //Add a name() method to the interface to show which payment method is being used
}

type CreditCard struct {
	transactionFeePercentage float64
}

type PayPal struct {
	processingFee float64
}

func (cc CreditCard) calculateFee(amount float64) float64 {
	return amount * cc.transactionFeePercentage
}

func (cc CreditCard) name() string {
	return "Credit Card"
}

func (pp PayPal) calculateFee(amount float64) float64 {
	return pp.processingFee
}

func (pp PayPal) name() string {
	return "PayPal"
}

func getProcessingFee(p PaymentMethod, amount float64) {
	fee := p.calculateFee(amount)
	fmt.Printf("Processing a ₹%.2f payment via %s. Fee: ₹%.2f\n", amount, p.name(), fee)
}

func main() {
	fmt.Println("code logic here")
	creditcard := CreditCard{transactionFeePercentage: 0.03}
	paypal := PayPal{processingFee: 0.04}

	getProcessingFee(creditcard, 111.11)
	getProcessingFee(paypal, 111.11)
}
