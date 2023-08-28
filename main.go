package main

import (
	"fmt"
)

func main() {
	calculator := Calculator{}

	a := 10.0
	b := 5.0

	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, calculator.Add(a, b))
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, calculator.Multiply(a, b))

	result, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, result)
	}
}
