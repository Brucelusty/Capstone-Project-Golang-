package main

import "fmt"

func lesson10_Pointers() {
	fmt.Println("\n=== LESSON 10: POINTERS ===")
	
	// Regular variable
	x := 42
	fmt.Printf("x = %d\n", x)
	
	// Pointer to x
	var ptr *int = &x // & gets the address
	fmt.Printf("Address of x: %p\n", ptr)
	fmt.Printf("Value at address (dereferencing): %d\n", *ptr) // * gets the value
	
	// Modify through pointer
	*ptr = 100
	fmt.Printf("x after pointer modification: %d\n", x)
	
	// Why use pointers?
	// 1. Avoid copying large structs
	// 2. Allow functions to modify the original value
	
	type BigStruct struct {
		Data [1000]int
	}
	
	// Function with pointer - efficient and can modify
	modifyValue := func(p *int) {
		*p = *p * 2
	}
	
	value := 50
	fmt.Printf("Before: %d\n", value)
	modifyValue(&value)
	fmt.Printf("After: %d\n", value)
	
	// nil pointer
	var nilPtr *int
	fmt.Printf("Nil pointer: %v\n", nilPtr)
	
	// Creating pointer with new
	ptr2 := new(int) // Allocates memory and returns pointer
	*ptr2 = 99
	fmt.Printf("Pointer created with new: %d\n", *ptr2)
}