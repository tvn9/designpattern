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
	// result := make([]*Product, 0)
	result := []*Product{}
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

// program entry
func main() {
	// create some products
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	car := Product{"Car", red, large}
	boat := Product{"Boat", green, medium}
	// Add proudct to slice of products
	products := []Product{apple, tree, car, boat}

	fmt.Println(products)
	fmt.Println("Green products (basic):")
	f := Filter{}
	for _, v := range f.FilterByClolor(products, green) {
		fmt.Printf(" - %s green\n", v.name)
	}

	fmt.Println("Medium size products")
	for _, v := range f.FilterbySize(products, medium) {
		fmt.Printf(" - %s medium\n", v.name)
	}

	fmt.Println("Red and Large Products")
	for _, v := range f.FilterBySizeColor(products, large, red) {
		fmt.Printf(" - %s large, red\n", v.name)
	}
}
