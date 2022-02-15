package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)

type PersonBuilder struct {
	action []personMod
}

func (b *PersonBuilder) Named(name string) *PersonBuilder {
	b.action = append(b.action, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Titled(title string) *PersonBuilder {
	b.action = append(b.action, func(p *Person) {
		p.position = title
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.action {
		a(&p)
	}
	return &p
}

func main() {

	// make a PersonBuilder
	b := PersonBuilder{}

	p := b.Named("Marko").Titled("Principle Software Engineer").Build()
	fmt.Println(*p)
}
