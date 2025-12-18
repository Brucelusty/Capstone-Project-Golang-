
// ============================================================================
// LESSON 9: ERROR HANDLING
// ============================================================================
package main

import "fmt"

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "cannot be negative",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Message: "is unrealistically high",
		}
	}
	return nil
}

func lesson9_ErrorHandling() {
	fmt.Println("\n=== LESSON 9: ERROR HANDLING ===")
	
	// Basic error handling
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
	
	// Error case
	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	
	// Custom error
	if err := validateAge(-5); err != nil {
		fmt.Println("Validation error:", err)
	}
	
	if err := validateAge(200); err != nil {
		fmt.Println("Validation error:", err)
	}
	
	if err := validateAge(25); err == nil {
		fmt.Println("Age 25 is valid!")
	}
}