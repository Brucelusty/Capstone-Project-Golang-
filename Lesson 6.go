package main

import "fmt"

func lesson6_Maps() {
	fmt.Println("\n=== LESSON 6: MAPS ===")
	
	// Creating a map
	var ages map[string]int // nil map - cannot be used yet
	ages = make(map[string]int) // Initialize
	
	// Adding key-value pairs
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	
	// Map literal
	grades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "A",
	}
	
	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("Grades: %v\n", grades)
	
	// Accessing values
	fmt.Printf("Alice's age: %d\n", ages["Alice"])
	fmt.Printf("Bob's grade: %s\n", grades["Bob"])
	
	// Check if key exists
	age, exists := ages["David"]
	if exists {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Println("David not found in ages map")
	}
	
	// Delete key
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)
	
	// Iterate over map
	fmt.Println("\nAll grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %s\n", name, grade)
	}
	
	// Map length
	fmt.Printf("Number of students: %d\n", len(grades))
}