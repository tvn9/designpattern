package main

import "fmt"

type Person struct {
	// address
	StreetAddress, City, State, Postcode string
	// job
	CompanyName, Position string
	AnnualIncome          int
}

// PersonBuilder struct
type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (pjb *PersonJobBuilder) At(compName string) *PersonJobBuilder {
	pjb.person.CompanyName = compName
	return pjb
}

func (pjb *PersonJobBuilder) Earnings(annual int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annual
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}
func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) State(state string) *PersonAddressBuilder {
	it.person.State = state
	return it
}

func (it *PersonAddressBuilder) WithPostcode(zipcode string) *PersonAddressBuilder {
	it.person.Postcode = zipcode
	return it
}

// NewPersonBuilder function initialization
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func main() {
	// Call NewPersonBuilder
	pb := NewPersonBuilder()
	pb.Lives().At("650 Harry Rd").In("San Jose, CA").State("California").WithPostcode("95120").
		Works().At("IBM Research - Almaden").AsA("Programmer").Earnings(150000)

	// Get the person information
	person := pb.Build()
	fmt.Println(*person)
}
