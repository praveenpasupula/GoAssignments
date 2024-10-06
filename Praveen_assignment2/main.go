// File: main.go
package main

import (
	"assignment2/calculator"
	"fmt"
)

func main() {
	a, b := 10.0, 5.0

	fmt.Printf("Addition: %.2f\n", calculator.Add(a, b))
	fmt.Printf("Subtraction: %.2f\n", calculator.Subtract(a, b))
	fmt.Printf("Multiplication: %.2f\n", calculator.Multiply(a, b))
	fmt.Printf("Division: %.2f\n", calculator.Divide(a, b))
}
