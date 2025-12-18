package main

import "fmt"

func lesson2_DataTypes() {
	fmt.Println("\n=== LESSON 2: DATA TYPES ===")
	
	// INTEGER TYPES
	var smallInt int8 = 127        // -128 to 127
	var mediumInt int16 = 32767    // -32,768 to 32,767
	var normalInt int32 = 2147483647
	var bigInt int64 = 9223372036854775807
	var regularInt int = 42        // Size depends on system (32 or 64 bit)
	
	// UNSIGNED INTEGERS (positive only)
	var positiveInt uint = 100
	var byteValue byte = 255       // Same as uint8 (0 to 255)
	
	// FLOATING POINT
	var price float32 = 19.99
	_ = price // Use the variable to avoid error
	
	// STRING
	var message string = "Hello, Go!"
	var multiline string = `This is a
	multi-line string
	using backticks`
	
	// BOOLEAN
	var isActive bool = true
	var isComplete bool = false
	
	// Type conversion (explicit in Go - no automatic conversion!)
	var intValue int = 42
	var floatValue float64 = float64(intValue) // Must explicitly convert
	var stringValue string = fmt.Sprintf("%d", intValue) // Convert to string
	
	fmt.Printf("Integer: %d, Float: %.2f, String: %s\n", 
		regularInt, price, message)
	fmt.Printf("Converted: int=%d, float=%.2f, string=%s\n", 
		intValue, floatValue, stringValue)
	fmt.Printf("Boolean: %t\n", isActive && !isComplete)
	fmt.Println(multiline)
	
	// Demonstrate all integer types
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		smallInt, mediumInt, normalInt, bigInt)
	fmt.Printf("uint: %d, byte: %d\n", positiveInt, byteValue)
}