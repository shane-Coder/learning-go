package main

import (
	"fmt"
	"math"
)

// 1. Define the interface (the contract)
// It says any type is a "Shape" if it has an area() method.
type Shape interface {
	area() float64
}

// 2. Create our custom types (structs)
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

// 3. Implement the interface for Circle
// Because Circle has an area() method, it now satisfies the Shape interface.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 4. Implement the interface for Rectangle
// Because Rectangle has an area() method, it also satisfies the Shape interface.
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// 5. Create a generic function that can work with ANY Shape
// This function doesn't care if it gets a Circle or a Rectangle.
// It only cares that the type passed to it has an area() method.
func printShapeInfo(s Shape) {
	fmt.Printf("The area of the shape is %0.2f\n", s.area())
}

func main() {
	// Create instances of our structs
	circle := Circle{radius: 8.33}
	rectangle := Rectangle{length: 11.11, width: 4.13}

	// Pass both types to the same function!
	printShapeInfo(circle)
	printShapeInfo(rectangle)
}
