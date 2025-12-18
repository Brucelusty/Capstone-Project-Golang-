

package main

import (
	"fmt"
)

// ============================================================================
// LESSON 1: BASIC SYNTAX AND VARIABLES
// ============================================================================


var globalVariable string = "I'm accessible everywhere" // Package-level variable

func lesson1_Variables() {
	fmt.Println("\n=== LESSON 1: VARIABLES ===")
	
	// Method 1: Using var with type
	var name string = "Bruce"
	var age int = 25
	var isStudent bool = true
	
	// Method 2: Type inference (Go figures out the type)
	var city = "Amsterdam"
	
	// Method 3: Short declaration (most common, only inside functions)
	country := "Netherlands"
	
	// Multiple declarations
	var x, y, z int = 1, 2, 3
	
	// Constants (cannot be changed)
	const pi = 3.14159
	const greeting = "Hello, World!"
	
	// Print all variables
	fmt.Printf("Name: %s, Age: %d, Student: %t\n", name, age, isStudent)
	fmt.Printf("Location: %s, %s\n", city, country)
	fmt.Printf("Numbers: %d, %d, %d\n", x, y, z)
	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Println(greeting)
	
	// IMPORTANT: Unused variables cause compile errors in Go!
	// This encourages clean code
}










