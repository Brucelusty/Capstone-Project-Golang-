package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func lesson8_Interfaces() {
	fmt.Println("\n=== LESSON 8: INTERFACES ===")
	
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	
	fmt.Println("Rectangle:")
	printShapeInfo(rect)
	
	fmt.Println("\nCircle:")
	printShapeInfo(circle)
	
	// Empty interface (interface{}) accepts any type
	var anything interface{}
	anything = "Hello"
	fmt.Printf("anything as string: %v\n", anything)
	anything = 42
	fmt.Printf("anything as int: %v\n", anything)
	anything = true
	fmt.Printf("anything as bool: %v\n", anything)
	
	// Type assertion
	var value interface{} = "Go is awesome"
	str, ok := value.(string)
	if ok {
		fmt.Printf("Type assertion successful: %s\n", str)
	}
}