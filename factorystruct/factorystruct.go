// Factory Struct approach
package main

import "fmt"

// Employee Struct
type Employee struct {
	Name, Position string
	Salary         int
}

// EmployeeFactory struct
type EmployeeFactory struct {
	Position string
	Salary   int
}

// Create method creates the Employee
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Salary}
}

// NewEmployeeFactory function create NewEmployeeFactory
func NewEmployeeFactory(position string, salary int) *EmployeeFactory {
	return &EmployeeFactory{position, salary}
}

// main entry
func main() {
	// make a developerFactory
	developerFactory := NewEmployeeFactory("Programmer", 200000)
	managerFactory := NewEmployeeFactory("Program Manager", 150000)

	// now make the developer
	developer := developerFactory.Create("Mike")
	manager := managerFactory.Create("Bob")

	// now print the infor
	fmt.Println(developer)
	fmt.Println(manager)
}
