package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstructions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)

// Info struct
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Person struct
type Person struct {
	name     string
	birthday string
	address  string
}

// Refactor RelationshipBrowser
type RelationshipBrowser interface {
	FindAllChild(name string) []*Person
}

// Relationships struct
type Relationships struct {
	relations []Info
}

// high-level module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationshipBrowser // refactor for DIP
}

func (r *Relationships) FindAllChild(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// AddPerentAndChild
func (r *Relationships) AddPerentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		from:         parent,
		relationship: Parent,
		to:           child,
	})
}

// Refactor Investigate func below
func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChild("John") {
		fmt.Println("John has a child called", p.name)
	}
}

// // Investigate
// func (r *Research) Investigate() {
// 	relations := r.relationships.relations // low level dependency, not recommanded
// 	for _, rel := range relations {
// 		if rel.from.name == "John" && rel.relationship == Parent {
// 			fmt.Println("John has a child called", rel.to.name)
// 		}
// 	}
// }

// program entry
func main() {
	// populate the Person struct
	parent := Person{"John", "08/19/1999", "2345 St. Patrick Rd, San Jose, CA 93122"}
	child1 := Person{"Chris", "06/30/1996", "883 St. John Rd, San Jose, CA 95123"}
	child2 := Person{"Matt", "05/16/1994", "556 St. Matthew Rd, San Jose, CA 95123"}
	child3 := Person{"Dave", "", ""}

	// Perform research into the relationships
	relationships := Relationships{}
	relationships.AddPerentAndChild(&parent, &child1)
	relationships.AddPerentAndChild(&parent, &child2)
	relationships.AddPerentAndChild(&parent, &child3)

	r := Research{&relationships}
	r.Investigate()

	//

}
