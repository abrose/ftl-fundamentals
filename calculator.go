// Package calculator provides a library for simple calculations in Go.
package calculator

import "fmt"
import "math"

// Add takes two numbers and returns the result of adding them together.
func Add(inputs ...float64) float64 {
	total := 0.0
	for _, input := range inputs {
		total += input
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(inputs ...float64) float64 {
	total := inputs[0]
	for _, input := range inputs[1:] {
		total -= input
	}
	return total
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

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input, %f (Square root of negative)", a)
	}
	return math.Sqrt(a), nil
}

// This function takes a string expression of the format %f (+|-|/|*) %f and evaluates it
func Evaluate(expression string) (float64, error) {
	var operandOne, operandTwo float64
	var operator string
	var err error

	_, err = fmt.Sscanf(expression, "%f %s %f", &operandOne, &operator, &operandTwo)

	if err != nil {
		return 0, err
	}

	var result float64

	switch operator {
	case "+":
		result = Add(operandOne, operandTwo)
	case "-":
		result = Subtract(operandOne, operandTwo)
	case "/":
		result, err = Divide(operandOne, operandTwo)
	case "*":
		result = Multiply(operandOne, operandTwo)
	default:
		result = 0
		err = fmt.Errorf("bad input, can not evaluate (%s)", expression)
	}

	return result, err
}
