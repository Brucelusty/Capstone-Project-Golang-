package main

import (
	"fmt"
	"strings"
)

func lesson13_Strings() {
	fmt.Println("\n=== LESSON 13: STRING OPERATIONS ===")
	
	text := "Hello, World!"
	
	// String length
	fmt.Printf("Length: %d\n", len(text))
	
	// Contains
	fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))
	
	// Replace
	newText := strings.Replace(text, "World", "Go", 1)
	fmt.Printf("Replaced: %s\n", newText)
	
	// Split
	words := strings.Split("apple,banana,orange", ",")
	fmt.Printf("Split: %v\n", words)
	
	// Join
	joined := strings.Join(words, " - ")
	fmt.Printf("Joined: %s\n", joined)
	
	// Upper/Lower case
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	
	// Trim
	trimmed := strings.TrimSpace("  Hello  ")
	fmt.Printf("Trimmed: '%s'\n", trimmed)
	
	// HasPrefix/HasSuffix
	fmt.Printf("Starts with 'Hello': %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("Ends with '!': %t\n", strings.HasSuffix(text, "!"))
}
