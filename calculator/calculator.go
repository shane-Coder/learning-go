// This file is part of the 'calculator' package.
package calculator

// 'Add' starts with a capital letter, so it is exported (public).
func Add(a, b int) int {
	return a + b
}

// 'subtract' starts with a lowercase letter, so it is unexported (private).
func subtract(a, b int) int {
	return a - b
}
