
// ============================================================================
// LESSON 11: GOROUTINES (CONCURRENCY)
// ============================================================================
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(150 * time.Millisecond)
	}
}

func lesson11_Goroutines() {
	fmt.Println("\n=== LESSON 11: GOROUTINES (CONCURRENCY) ===")
	
	// Run functions concurrently with 'go' keyword
	go printNumbers()
	go printLetters()
	
	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
	
	fmt.Println("\nAll goroutines completed!")
}
