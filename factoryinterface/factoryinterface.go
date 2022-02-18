// Interface factory
package main

import "fmt"

// define an interface
type Person interface {
	SayHello()
}

// define a person struct
type person struct {
	name string
	age  int
}

type oldPerson struct {
	person person
}

//
func (p *person) SayHello() {
	fmt.Printf("Hello my name is %s, I am %d years young, let go hiking!\n", p.name, p.age)
}

func (p *oldPerson) SayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old, sorry I am too old to go hiking with you!\n",
		p.person.name, p.person.age)
}

// Define a function contructor
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &oldPerson{person{name, age}}
	} else {
		return &person{name, age}
	}
}

// program entry
func main() {
	// make a person
	mark := NewPerson("Mark", 38)
	mark.SayHello()
	john := NewPerson("John", 109)
	john.SayHello()

}
