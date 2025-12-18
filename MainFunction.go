
// ============================================================================
// MAIN FUNCTION - RUN ALL LESSONS
// ============================================================================
package main

import "fmt"

func main() {
	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Println("║   GO LANGUAGE COMPLETE BEGINNER'S GUIDE             ║")
	fmt.Println("║   Learning Go from Basics to Advanced Concepts      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
	
	// Run all lessons
	lesson1_Variables()
	lesson2_DataTypes()
	lesson3_ControlStructures()
	lesson4_Functions()
	lesson5_ArraysAndSlices()
	lesson6_Maps()
	lesson7_Structs()
	lesson8_Interfaces()
	lesson9_ErrorHandling()
	lesson10_Pointers()
	lesson11_Goroutines()
	lesson12_Channels()
	lesson13_Strings()
	
	fmt.Println("\n╔══════════════════════════════════════════════════════╗")
	fmt.Println("║   🎉 CONGRATULATIONS!                               ║")
	fmt.Println("║   You've completed all Go beginner lessons!         ║")
	fmt.Println("║                                                      ║")
	fmt.Println("║   Next Steps:                                        ║")
	fmt.Println("║   • Build a REST API                                 ║")
	fmt.Println("║   • Connect to a database                            ║")
	fmt.Println("║   • Create a CLI tool                                ║")
	fmt.Println("║   • Explore Go web frameworks (Gin, Echo)            ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
}