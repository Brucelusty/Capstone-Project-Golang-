
// ============================================================================
// LESSON 4: FUNCTIONS
// ============================================================================
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from a function!")
}

// Function with parameters
func greet(name string, age int) {
	fmt.Printf("Hello %s, you are %d years old\n", name, age)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values (very common in Go!)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // Returns named variables automatically
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as a parameter (higher-order function)
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func lesson4_Functions() {
	fmt.Println("\n=== LESSON 4: FUNCTIONS ===")
	
	sayHello()
	greet("Alice", 30)
	
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}
	
	// Named returns
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", area, perimeter)
	
	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	
	// Function as parameter
	multiply := func(x, y int) int { return x * y }
	result2 := applyOperation(4, 5, multiply)
	fmt.Printf("4 * 5 = %d\n", result2)
}