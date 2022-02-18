package main

import "fmt"

// Deep Copying Pattern

type Address struct {
	StreetAddress, City, State, ZipCode, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
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

//
func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
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
		Friends: []string{"Chris", "Tim", "Mike"},
	}

	chris := mark.DeepCopy()
	chris.Name = "Chris"
	tim := mark.DeepCopy()
	tim.Name = "Tim"

	chris.Address.StreetAddress = "999 Sunset Blvd"

	fmt.Println(mark, mark.Address)
	fmt.Println(chris, chris.Address)
	fmt.Println(tim, tim.Address)

	chris.Address = &Address{
		StreetAddress: "555 St.",
		City:          "Sun City",
		State:         "Ok",
		ZipCode:       "33333",
		Country:       "USA",
	}

	chris.Friends = append(chris.Friends, "Chris", "Time", "Marko")
	tim.Address = &Address{"666 Ave", "Daily City", "Ca", "95322", "USA"}
	tim.Friends = append(tim.Friends, "Chris", "Time", "Mike")

	fmt.Println(mark, chris, tim)
}
