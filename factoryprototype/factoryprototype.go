// Factory prototype pattern
package main

import (
	"fmt"
	"os"
)

// Employee struct
type Employee struct {
	Name, Position string
	Salary         int
}

// const setup option for switch case
const (
	Developer = iota // option 0
	Manager          // option 1
)

// NewEmployee function for creating new NewEmployee
func NewEmployee(role int) *Employee {
	switch role {
	case Developer: // option 0
		return &Employee{"", "Developer", 160000}
	case Manager: // option 1
		return &Employee{"", "Manager", 200000}
	default:
		fmt.Fprintln(os.Stderr, "Unsupported role")
	}
	return &Employee{}
}

// main entry
func main() {
	// make en Employee
	m := NewEmployee(Manager)
	m.Name = "Mike"

	v := NewEmployee(Developer)
	v.Name = "Victor"
	// Print the result
	fmt.Println(m)
	fmt.Println(v)
}
