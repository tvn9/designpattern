package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

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
// low-level module
type RelationshipBrowser interface {
	FindAllChildOf(name string) []*Person
}

// Relationships struct
type Relationships struct {
	// if low level storage has changed, then the code from highlevel will break
	relations []Info
}

// FindAllChildOf - refactor
func (r *Relationships) FindAllChildOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// high-level module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationshipBrowser // refactor for DIP
}

// AddPerentAndChildOf
func (r *Relationships) AddPerentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{from: parent, relationship: Parent, to: child})
	r.relations = append(r.relations, Info{from: child, relationship: Child, to: parent})
}

// Refactor Investigate func below
func (r *Research) Investigate() {
	for _, rel := range r.browser.FindAllChildOf("John") {
		fmt.Println("John has a child called", rel.name)
	}
}

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
