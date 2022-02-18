// Factory function example
package main

import "fmt"

// define a Person struct
type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// NewPerson factory function
func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2}
}

// program entry
func main() {
	// initialize a Person
	p := NewPerson("John", 33)
	//
	fmt.Printf("Hello I am %s, I am %d year young\n", p.Name, p.Age)
}
