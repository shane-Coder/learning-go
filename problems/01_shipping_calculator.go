/*
Problem Statement #1: Shipping Cost Calculator

Objective:
	Write a Go program that calculates the total shipping cost for a package based on its destination, weight, and priority.

Rules/Logic:

1.	Base Cost (use a switch statement for this):

		If the destinationZone is "Domestic", the base cost is $10.

		If the destinationZone is "International", the base cost is $25.

		For any other zone, the base cost is $50.

2.	Surcharges (use if/else if/else for this):

		If the packageWeight (in kg) is over 20kg, add a $10 heavy package surcharge to the total cost.

		If the isPriority shipping is true:

			AND the destinationZone is "Domestic", add a $5 express surcharge.

			AND the destinationZone is "International", add a $15 express surcharge.

---
Terminal Output:

	// Test Case 1: Domestic, standard weight, no priority
	DestinationZone: Domestic, PackageWeight: 15, IsPriority: false
	The total shipping cost is: $10.00

	// Test Case 2: International, heavy weight, with priority
	DestinationZone: International, PackageWeight: 45, IsPriority: true
	The total shipping cost is: $50.00

	// Test Case 3: Default zone, heavy weight
	DestinationZone: Space, PackageWeight: 65, IsPriority: false
	The total shipping cost is: $60.00
*/

package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Shipping Cost Calculator!")

	destinationZone := "Domestic"
	packageWeight := 25.0
	isPriority := true
	baseCost := 0.0
	surcharge := 0.0

	fmt.Println("DestinationZone:", destinationZone, ", PackageWeight:", packageWeight, ", IsPriority:", isPriority)

	// creating task 1
	switch destinationZone {
	case "Domestic":
		baseCost += 10
	case "International":
		baseCost += 25
	default:
		baseCost += 50
	}

	// creating task 2
	if packageWeight > 20 {
		surcharge += 10
	}
	if isPriority && destinationZone == "Domestic" {
		surcharge += 5
	} else if isPriority && destinationZone == "International" {
		surcharge += 15
	}

	// total shipping cost
	totalCost := baseCost + surcharge
	fmt.Printf("The total shipping cost is: $%.2f\n", totalCost)
}
