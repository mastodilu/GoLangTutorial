// Package textanalyze helps counting the number of occurrences for each word in a text.
package textanalyze

import (
	"fmt"
	"strings"
)

// Wordsmap is a map that counts the number of occurrences for each word in the text
var Wordsmap map[string]int

// CountWords counts each words using a map.
func CountWords(text string) {
	Wordsmap = make(map[string]int)
	sslice := strings.Fields(text)
	for _, word := range sslice {
		Wordsmap[word]++
	}
}

// PrintMap prints the whole map in a readable way
func PrintMap() {
	for k, v := range Wordsmap {
		fmt.Println(k, v)
	}
}
