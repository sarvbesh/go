package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each "word" in the string s
func WordCount(s string) map[string]int {
	
	// initialize the map to store the word counts
	counts := make(map[string]int)
	
	// split the input string into a slice of words using strings.Fields
	// strings.Fields splits the string around any white space (spaces, tabs, newlines)
	words := strings.Fields(s)
	
	// 3. iterate over the slice of words
	for _, word := range words {

		// increment the count for the current word in the map
		counts[word]++
	}
	
	return counts
}

func main() {
	wc.Test(WordCount)
}