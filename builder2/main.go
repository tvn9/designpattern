package main

import (
	"fmt"
	"strings"
)

const (
	indent = 2
)

type HTMLElement struct {
	tag, text string
	elements  []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(in int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indent*in)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.tag))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indent*(in+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(in + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.tag))
	return sb.String()
}

// HTMLBuilder
type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

// NewHTMLBuilder
func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName, HTMLElement{rootName, "", []HTMLElement{}}}
}

//
func (b *HTMLBuilder) String() string {
	return b.root.String()
}

//
func (b *HTMLBuilder) AddChild(chName, chText string) {
	e := HTMLElement{tag: chName, text: chText, elements: []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

// Program entry
func main() {
	nb := NewHTMLBuilder("html")
	nb.AddChild("head", "")
	nb.AddChild("body", "")
	nb.AddChild("h1", "Hello Builder Design Pattern")
	nb.AddChild("p", "Design pattern is a mazing and fund to learn")
	fmt.Println(nb.String())
}
