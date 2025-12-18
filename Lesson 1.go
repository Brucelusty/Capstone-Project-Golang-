// ============================================================================
// GO LANGUAGE BEGINNER'S GUIDE WITH COMPREHENSIVE DOCUMENTATION
// ============================================================================
// Author: Bruce
// Purpose: Complete guide for learning Go from scratch
// Topics: Variables, Functions, Structs, Interfaces, Concurrency, and more
// ============================================================================

package main

import (
	"fmt"
)

// ============================================================================
// LESSON 1: BASIC SYNTAX AND VARIABLES
// ============================================================================

// Variables can be declared in multiple ways
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

// ============================================================================
// LESSON 2: DATA TYPES
// ============================================================================



// ============================================================================
// LESSON 3: CONTROL STRUCTURES (IF, SWITCH, LOOPS)
// ============================================================================



// ============================================================================
// LESSON 4: FUNCTIONS
// ============================================================================

// Simple function with no parameters or return value


// ============================================================================
// LESSON 5: ARRAYS AND SLICES
// ============================================================================



// ============================================================================
// LESSON 6: MAPS (Key-Value Pairs)
// ============================================================================



// ============================================================================
// LESSON 7: STRUCTS (Custom Types)
// ============================================================================

// Define a struct (like a class in other languages)


// ============================================================================
// LESSON 8: INTERFACES
// ============================================================================

// Interface defines behavior (methods)


// ============================================================================
// LESSON 9: ERROR HANDLING
// ============================================================================

// Function that returns an error


// ============================================================================
// LESSON 10: POINTERS
// ============================================================================



// ============================================================================
// LESSON 11: GOROUTINES (CONCURRENCY)
// ============================================================================



// ============================================================================
// LESSON 12: CHANNELS (COMMUNICATION BETWEEN GOROUTINES)
// ============================================================================






// ============================================================================
// MAIN FUNCTION - RUN ALL LESSONS
// ============================================================================



// ============================================================================
// QUICK REFERENCE GUIDE
// ============================================================================
/*

BASIC SYNTAX:
- Variables: var name type = value  OR  name := value
- Constants: const name = value
- Functions: func name(params) returnType { }
- Comments: // single line  OR  /* multi-line */

/*
DATA TYPES:
- int, int8, int16, int32, int64
- uint, uint8 (byte), uint16, uint32, uint64
- float32, float64
- string
- bool
- rune (Unicode character)

OPERATORS:
- Arithmetic: + - * / % ++ --
- Comparison: == != < > <= >=
- Logical: && || !
- Bitwise: & | ^ << >>

CONTROL FLOW:
- if/else if/else
- switch
- for (only loop in Go)
- break, continue

COLLECTIONS:
- Arrays: [size]type
- Slices: []type
- Maps: map[keyType]valueType

ADVANCED:
- Structs: Custom types
- Interfaces: Define behavior
- Pointers: &variable, *pointer
- Goroutines: go function()
- Channels: make(chan type)

ERROR HANDLING:
- Always check errors: if err != nil { }
- Return errors from functions
- Use custom error types

BEST PRACTICES:
1. Use short variable names in small scopes
2. Always handle errors
3. Use gofmt to format code
4. Write comments for exported functions
5. Keep functions small and focused
6. Use interfaces for flexibility
7. Prefer composition over inheritance

NAMING CONVENTIONS:
- Exported: Start with uppercase (Public)
- Unexported: Start with lowercase (Private)
- camelCase for multi-word names
- MixedCaps for exported multi-word names

*/