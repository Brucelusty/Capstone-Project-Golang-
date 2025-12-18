
// ============================================================================
// LESSON 7: STRUCTS (Custom Types)
// ============================================================================
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

// Method on struct (function associated with Person)
func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

// Struct with embedded struct (composition)
type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person  // Embedded struct
	Address // Embedded struct
	JobTitle string
	Salary   float64
}

func lesson7_Structs() {
	fmt.Println("\n=== LESSON 7: STRUCTS ===")
	
	// Creating structs
	person1 := Person{
		Name:    "Alice",
		Age:     25,
		Email:   "alice@example.com",
		IsAdmin: false,
	}
	
	// Shorthand (order matters)
	person2 := Person{"Bob", 30, "bob@example.com", true}
	
	// Partial initialization
	person3 := Person{Name: "Charlie", Age: 35}
	
	fmt.Printf("Person 1: %+v\n", person1) // %+v shows field names
	fmt.Printf("Person 2: %v\n", person2)
	fmt.Printf("Person 3: %v\n", person3)
	
	// Accessing fields
	fmt.Printf("%s's email: %s\n", person1.Name, person1.Email)
	
	// Calling methods
	person1.Greet()
	person1.HaveBirthday()
	
	// Embedded struct
	employee := Employee{
		Person: Person{
			Name:  "David",
			Age:   28,
			Email: "david@company.com",
		},
		Address: Address{
			Street:  "Main St 123",
			City:    "Amsterdam",
			Country: "Netherlands",
		},
		JobTitle: "Software Engineer",
		Salary:   75000.00,
	}
	
	// Access embedded fields directly
	fmt.Printf("\nEmployee: %s\n", employee.Name) // From Person
	fmt.Printf("City: %s\n", employee.City)       // From Address
	fmt.Printf("Job: %s, Salary: $%.2f\n", employee.JobTitle, employee.Salary)
}