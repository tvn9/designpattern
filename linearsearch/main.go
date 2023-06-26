// Linearsearch
package main

import "fmt"

func linearsearch(list []string, word string) bool {
	if len(list) > 0 {
		for _, w := range list {
			if w == word {
				return true
			}
		}
	}
	return false
}

func main() {
	words := []string{"Apple", "Mango", "Banana", "Rasberry"}
	search := "Mango"

	fmt.Printf("Looking for %s return %v\n", search, linearsearch(words, search))
}
