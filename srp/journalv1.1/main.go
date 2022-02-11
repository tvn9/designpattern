package main

import (
	"dpgo/srp/journalv1.1/journal"
	"fmt"
	"os"
)

// program entry
func main() {
	j := journal.Journal{}
	j.AddEntry("Today is the firtday at work, and I am excited")
	j.AddEntry("The second day at work is also fun and productive")
	j.AddEntry("The third day our team when to team building activities event")
	fmt.Println("Before delete")
	for _, t := range j.Entries {
		fmt.Println(t)
	}

	if err := j.DelEntry(2); err != nil {
		os.Exit(1)
	}

	j.Save("journal.txt")
	fmt.Println("After delete")
	for _, t := range j.Entries {
		fmt.Println(t)
	}
}
