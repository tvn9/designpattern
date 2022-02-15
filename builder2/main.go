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

// AddChildFluent
func (b *HTMLBuilder) AddChildFluent(chName, chText string) *HTMLBuilder {
	e := HTMLElement{tag: chName, text: chText, elements: []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

// Program entry
func main() {
	nb := NewHTMLBuilder("html")
	nb.AddChild("head", "")
	nb.AddChild("body", "<h1>Hello Builder Design Patter</h1><p>Design patter is amazin and fun to learn</p>")
	fmt.Println(nb.String())

	ul := NewHTMLBuilder("html")
	ul.AddChild("head", "")
	ul.AddChildFluent("body", "<h1>Welcome to Build Patter!</h1>").
		AddChildFluent("ul", "").AddChildFluent("li", "Item 1").
		AddChildFluent("li", "Item 2").AddChildFluent("li", "Item 3").
		AddChildFluent("li", "Item 4")
	fmt.Println(ul.String())
}
