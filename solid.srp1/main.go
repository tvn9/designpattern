package main

import (
	"fmt"
	"os"
	"strings"
)

/////////////////////// Single Responsibility Princile /////////////////////
// The SRP recommanded that a software component should focus on a single
// responsibility where the featue impleted are only serve the main function
// of the feature. In this example the journal must only focus on:
// 1. Create journal record
// 2. Remove journal record
// 3. Keeping track of a list of journal entries
// In this case the journal entries are only exist in system memory while the
// App is running, persitent record should not be a responsilbity of this
// implementation.

// entryCount hold entry counter from 1++...
var entryCount = 0

// Journal struct holds a slice of Journal's entries
type Journal struct {
	entries []string
}

// implements Stringer interface
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// AddEntry adds an journal item to the list
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// DelEntry delete an jouritem from the list
func (j *Journal) DelEntry(in int) error {
	jl := j.entries
	if in <= 0 || in > len(jl) {
		return fmt.Errorf("journal entries %d not exist", in)
	}
	// Remove Journal entry
	j.entries = append(j.entries[:in-1], j.entries[in:]...)
	return nil
}

// program entry
func main() {
	myJournal := Journal{}
	myJournal.AddEntry("Today is the firtday at work, and I am excited")
	myJournal.AddEntry("The second day at work is also fun and productive")
	for _, t := range myJournal.entries {
		fmt.Println(t)
	}

	if err := myJournal.DelEntry(2); err != nil {
		os.Exit(1)
	}
	fmt.Println(myJournal.String())
}
