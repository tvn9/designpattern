package main

import "fmt"

// Deep Copying Pattern

type Address struct {
	StreetAddress, City, State, ZipCode, Country string
}

type Person struct {
	Name    string
	Address *Address
}

//
func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.State,
		a.ZipCode,
		a.Country,
	}
}

func main() {
	// make a person mark
	mark := &Person{
		Name: "Mark",
		Address: &Address{
			StreetAddress: "1255 Clay St.",
			City:          "San Jose",
			State:         "Ca",
			ZipCode:       "95122",
			Country:       "USA",
		},
	}

	chris := mark
	chris.Name = "Chris"

	chris.Address = &Address{
		mark.Address.StreetAddress,
		mark.Address.City,
		mark.Address.State,
		mark.Address.ZipCode,
		mark.Address.Country,
	}

	chris.Address.StreetAddress = "999 Sunset Blvd"
	fmt.Println(mark, mark.Address)
	fmt.Println(chris, chris.Address)

	chris.Address = &Address{
		StreetAddress: "555 St.",
		City:          "Sun City",
		State:         "Ok",
		ZipCode:       "33333",
		Country:       "USA",
	}

	fmt.Println(mark, chris)
}
