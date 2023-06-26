package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++
}

func (p Person) DoesNotGrow() {
	p.Age++
}

func main() {
	thanh := Person{"Person A", 10}
	thanh.Grow()

	fmt.Println("Person A's age:", thanh.Name, thanh.Age)

	ptr := &thanh
	ptr.DoesNotGrow()

	fmt.Println("Person A's age:", thanh.Name, thanh.Age)

	ptr.Grow()
	fmt.Println(ptr.Age)
}
