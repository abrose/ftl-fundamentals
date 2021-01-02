// Package calculator provides a library for simple calculations in Go.
package calculator

import "fmt"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying the first with
// the second
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of the division of the firstn
// number by the second.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (Division by zero undefined)", a, b)
	} else {
		return a / b, nil
	}
}

func Sqrt(a float64) float64 {
	return 1
}

