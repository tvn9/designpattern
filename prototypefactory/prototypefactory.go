// Prototype factory pattern
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

var mainOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "667 Virginia St.",
		City:          "Oakland",
	},
}

var auxOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "999 Main St.",
		City:          "Los Angeless",
	},
}

// DeepCopy make copy object through serialization
func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("Mark", 100)
	chris := NewAuxOfficeEmployee("Chris", 200)

	fmt.Println(john)
	fmt.Println(chris)
}
