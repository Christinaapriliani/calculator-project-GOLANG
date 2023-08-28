package main

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Add(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, but got %f", result)
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Subtract(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, but got %f", result)
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Multiply(3, 7)
	if result != 21 {
		t.Errorf("Expected 21, but got %f", result)
	}
}

func TestCalculator_Divide(t *testing.T) {
	calculator := Calculator{}

	// Test valid division
	result, err := calculator.Divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, but got %f", result)
	}

	// Test division by zero
	_, err = calculator.Divide(8, 0)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
