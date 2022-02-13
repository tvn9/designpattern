package main

import "fmt"

// OCP - Type should be open for extension but close for modification

// Color defines Color as an int
type Color int

// Color = iota assign 0 to Color and increment by 1 down the color list
const (
	brown  Color = iota // so brown is type Color and has value of 0
	red                 // red has value or 0 + 1 = 1
	orange              // orange = 2
	yellow              // = 3
	green
	blue
	purple
	grep
	white
	black
)

// Size defines Size as an int
type Size int

// Size = iota assign 0 to first size in the list and
// increment by 1 down list
const (
	small  Size = iota // small is type Size with value of 0
	medium             // = 1
	large
	xlarge
	xxlarge
)

// Product structure hold basic Product name, color, and size
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter defines and empty Filter struct
type Filter struct {
	//
}

// FilterByClolor compares color with color in product slice
func (f *Filter) FilterByClolor(p []Product, c Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if v.color == c {
			result = append(result, &p[i])
		}
	}
	return result
}

// FilterbySize
func (f *Filter) FilterbySize(p []Product, s Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if v.size == s {
			result = append(result, &p[i])
		}
	}
	return result
}

// FilterBySizeColor (Keep adding additional function could violate the OCP principal)
func (f *Filter) FilterBySizeColor(p []Product, s Size, c Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if v.size == s && v.color == c {
			result = append(result, &p[i])
		}
	}
	return result
}

/////////// Specification pattern //////////////////
type Specification interface {
	IsSatisfied(p *Product) bool
}

// ColorSpec struct
type ColorSpec struct {
	color Color
}

// IsSatisfied compares ColorSpec field color with Product's
// color field and return a boolean value
func (c ColorSpec) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// SizeSpec tructure
type SizeSpec struct {
	size Size
}

// IsSatisfied method compare SizeSpec's size with Product's
// size field and return a boolean value
func (s SizeSpec) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// BatterFilter is and empty struct
type BetterFilter struct{}

// Filter method compares Specification's field to Product's fields
func (f *BetterFilter) Filter(p []Product, s Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range p {
		if s.IsSatisfied(&v) {
			result = append(result, &p[i])
		}
	}
	return result
}

/////////////////// Composite Specification pattern /////////////////////

// CombineSpec defines first and second fields that the Filter must
// match in order to satisfy the search
type CombineSpec struct {
	first, second Specification
}

// IsSatisfied compares first and second fields of the Product's
//
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
	greenSpec := ColorSpec{green} // greenSpec variable
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// Set the SizeSpec to large and CombineSpec's first = greenSpec, second = large
	lSpec := SizeSpec{large}
	lgSpec := CombineSpec{
		first:  greenSpec, // define online 131 or command // greenSpec variable
		second: lSpec,
	}
	// Print filter results from Filter method
	fmt.Println("Combine Large & Green products:")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s green\n", v.name)
	}
}
