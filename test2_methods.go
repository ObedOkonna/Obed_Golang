package main

import (
	"fmt"
)

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Method with a receiver that modifies the struct
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Method without a receiver that accesses struct data without modifying it
func GetPersonAge(p Person) int {
	return p.Age
}

func main() {
	// Creating an instance of the struct
	person := Person{Name: "Alice", Age: 30}

	// Modifying struct data using the method
	person.UpdateAge(35)

	// Accessing struct data without modifying it using the method
	age := GetPersonAge(person)

	// Displaying the modified age
	fmt.Println("Modified Age:", age) // Output: Modified Age: 35
}
