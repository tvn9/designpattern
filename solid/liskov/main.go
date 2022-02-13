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
	if w <= 0 {
		fmt.Printf("could not set width = %d\n", w)
		return
	}
	r.width = w
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(h int) {
	if h <= 0 {
		fmt.Printf("could not set height = %d\n", 0)
		return
	}
	r.height = h
}

//////////////////// Breaking the List of Substitution Principle //////////////////
type Square struct {
	width, height int
}

// This violates the list of subtitution of principle
func (s *Square) SetWidth(w int) {
	s.width = w  // This is where the problem comes in
	s.height = w // the code tried to enforce both the hight and width to the same value
}
func (s *Square) SetHeight(h int) {
	s.height = h
	s.width = h
}

// UseIt
func UseIt(s Sized) {
	s.SetWidth(10)
	s.SetHeight(20)
	w := s.GetWidth()
	h := s.GetHeight()
	expectedArea := w * h
	actualArea := s.GetWidth() * s.GetHeight()
	fmt.Println("Expected an area of", expectedArea, "but got", actualArea)
}

func main() {
	// make a Rectangle
	rec := &Rectangle{2, 3}
	UseIt(rec)
	//

	//
}
