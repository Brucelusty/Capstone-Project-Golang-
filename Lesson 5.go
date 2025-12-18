package main

import "fmt"

func lesson5_ArraysAndSlices() {
	fmt.Println("\n=== LESSON 5: ARRAYS AND SLICES ===")
	
	// ARRAYS - Fixed size, cannot be changed
	var numbers [5]int // Array of 5 integers, initialized to 0
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	
	// Array with initialization
	fruits := [3]string{"Apple", "Banana", "Orange"}
	
	fmt.Printf("Array: %v, Length: %d\n", numbers, len(numbers))
	fmt.Printf("Fruits: %v\n", fruits)
	
	// SLICES - Dynamic size, most commonly used
	var colors []string // Empty slice
	colors = append(colors, "Red")
	colors = append(colors, "Green", "Blue") // Can append multiple
	
	// Slice with make (preallocated)
	scores := make([]int, 3, 5) // length=3, capacity=5
	scores[0] = 85
	scores[1] = 90
	scores[2] = 78
	
	// Slice literal
	names := []string{"Alice", "Bob", "Charlie", "David"}
	
	fmt.Printf("Colors: %v, Length: %d\n", colors, len(colors))
	fmt.Printf("Scores: %v, Length: %d, Capacity: %d\n", 
		scores, len(scores), cap(scores))
	
	// Slice operations
	fmt.Printf("All names: %v\n", names)
	fmt.Printf("First 2: %v\n", names[0:2])
	fmt.Printf("From index 2: %v\n", names[2:])
	fmt.Printf("Up to index 3: %v\n", names[:3])
	
	// Range over slice
	fmt.Println("\nIterating with range:")
	for index, value := range names {
		fmt.Printf("Index %d: %s\n", index, value)
	}
	
	// Ignore index with _
	fmt.Println("\nJust values:")
	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}
	
	// 2D Slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("\nMatrix: %v\n", matrix)
}