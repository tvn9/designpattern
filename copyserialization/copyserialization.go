package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Deep Copying Pattern

type Address struct {
	StreetAddress, City, State, ZipCode, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// DeepCopy make copy object through serialization
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
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

	fmt.Println(mark, mark.Address)
	fmt.Println(chris, chris.Address)

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

	fmt.Println(chris, chris.Address)
	fmt.Println(tim, tim.Address)
}
