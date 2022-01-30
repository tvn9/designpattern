package main

import (
	"designpatterns/persit"
	"fmt"
)

var entryCount = 0

var jn = persit.Journal{}

func AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	jn.Entries = append(jn.Entries, entry)
	return entryCount
}

func RemoveEntry(index int) {
	// ....
}

func main() {
	AddEntry("Hello, Journal!")
	AddEntry("Testing SRP concept!")
	fmt.Println(jn)

	//
	persit.SaveToFile(&jn, "journal.txt")
}
