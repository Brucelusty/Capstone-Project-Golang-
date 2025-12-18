
// ============================================================================
// LESSON 3: CONTROL STRUCTURES (IF, SWITCH, LOOPS)
// ============================================================================
package main

import "fmt"

func lesson3_ControlStructures() {
	fmt.Println("\n=== LESSON 3: CONTROL STRUCTURES ===")
	
	// IF-ELSE STATEMENTS
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}
	
	// IF with initialization (variable only exists in if scope)
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}
	
	// SWITCH STATEMENTS (no break needed - auto breaks!)
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("Almost weekend!")
	case "Saturday", "Sunday": // Multiple cases
		fmt.Println("Weekend!")
	default:
		fmt.Println("Middle of the week")
	}
	
	// Switch with conditions (no variable)
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing!")
	case temperature < 15:
		fmt.Println("Cold")
	case temperature < 25:
		fmt.Println("Nice")
	default:
		fmt.Println("Hot!")
	}
	
	// FOR LOOPS (Go's only loop - no while or do-while)
	
	// Standard for loop
	fmt.Println("\nCounting 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	
	// While-style loop
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}
	
	// Infinite loop (use with break)
	attempts := 0
	for {
		attempts++
		if attempts > 3 {
			break // Exit loop
		}
		fmt.Printf("Attempt %d\n", attempts)
	}
	
	// Continue statement (skip iteration)
	fmt.Println("\nEven numbers from 1-10:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // Skip odd numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}