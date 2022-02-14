package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(w int)
	GetHeight() int
	SetHeight(h int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(w int) {
	r.width = w
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(h int) {
	r.height = h
}

func (r *Rectangle) NewRectangle(w, h int) *Rectangle {
	r.width = w
	r.height = h
	return &Rectangle{r.width, r.height}
}

//////////////////// Breaking the List of Substitution Principle //////////////////
type Square struct {
	Rectangle
}

// NewSquare
func NewSquare(l int) *Square {
	sq := Square{}
	sq.width = l
	sq.height = l
	return &sq
}

func (s *Square) GetWidth() int {
	return s.width
}
func (s *Square) SetWidth(w int) {
	s.width = w  // to enforce that the object is a square we set both width and height
	s.height = w // to the same value
}
func (s *Square) GetHeight() int {
	return s.height
}
func (s *Square) SetHeight(h int) {
	s.height = h // same here to enforce the share is the square we set
	s.width = h  // the width and height to same value
}

// UseIt
func UseIt(s Sized) {
	width := s.GetWidth()
	s.SetHeight(10)
	expectedArea := 10 * width
	actualArea := s.GetWidth() * s.GetHeight()
	fmt.Println("Expected an area of", expectedArea, "but got", actualArea)
}

func main() {
	// make a Rectangle
	rec := &Rectangle{2, 3}
	UseIt(rec)

	sq := NewSquare(5)
	UseIt(sq)
}
