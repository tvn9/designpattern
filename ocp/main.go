package main

import "fmt"

// OCP
// open for extension, close for modification
//

type Color int

const (
	brown Color = iota
	red
	orange
	yellow
	green
	blue
	purple
	grep
	white
	black
)

type Size int

const (
	small Size = iota
	medium
	large
	xlarge
	xxlarge
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) FilterByClolor(p []Product, c Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range p {
		if v.color == c {
			result = append(result, &p[i])
		}
	}
	return result
}

// Specification pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpec struct {
	color Color
}

func (c ColorSpec) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpec struct {
	size Size
}

func (s SizeSpec) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(p []Product, s Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if s.IsSatisfied(&v) {
			result = append(result, &p[i])
		}
	}
	return result
}

// Composite Specification pattern
type CombineSpec struct {
	first, second Specification
}

func (c CombineSpec) IsSatisfied(p *Product) bool {
	return c.first.IsSatisfied(p) && c.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	car := Product{"Car", red, large}
	//
	products := []Product{apple, tree, car}
	fmt.Println(products)
	fmt.Println("Green products (old):")
	f := Filter{}
	for _, v := range f.FilterByClolor(products, green) {
		fmt.Printf(" - %s green\n", v.name)
	}
	//
	fmt.Println("Green products (new):")
	greenSpec := ColorSpec{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	lSpec := SizeSpec{large}
	lgSpec := CombineSpec{greenSpec, lSpec}

	fmt.Println("Combine Large & Green products:")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s green\n", v.name)
	}

}
