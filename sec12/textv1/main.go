package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

// String
func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

// Capitalize method
func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

// NewFormattedText
func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText,
		make([]bool, len(plainText))}
}

// program entry
func main() {
	//
	text := "This is a weird new world"
	ft := NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())
}

// Example of an exteamly inefficent
