package main

import (
	"dpgo/solid/srp/journalv1.1/journal"
	"fmt"
	"os"
)

// program entry
func main() {
	j := journal.Journal{}
	j.AddEntry("Today is the firstday at work, and I am excited")
	j.AddEntry("The second day at work is also fun and productive")
	j.AddEntry("The third day our team when to team building event")
	fmt.Println("Before delete")
	// for _, t := range j.Entries {
	// 	fmt.Println(t)
	// }
	fmt.Println(j.String())

	// Delete entry 2
	if err := j.DelEntry(2); err != nil {
		os.Exit(1)
	}

	// SaveToFile saves the journal data in memory to a file
	journal.SaveToFile(&j, "journal1.txt")
	fmt.Println("After delete")
	// for _, t := range j.Entries {
	// 	fmt.Println(t)
	// }
	fmt.Println(j.String())
}
