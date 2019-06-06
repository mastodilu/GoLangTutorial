// Package textanalyze helps counting the number of occurrences for each word in a text.
package textanalyze

import (
	"fmt"
	"strings"
)

// Wordsmap is a map that counts the number of occurrences for each word in the text
var Wordsmap map[string]int

// wordsCount stores the total numbers of words.
var wordsCount int

// CountWords counts each words using a map.
func CountWords(text string) {
	Wordsmap = make(map[string]int)
	sslice := strings.Fields(text)

	wordsCount = len(sslice)

	for _, word := range sslice {
		Wordsmap[word]++
	}
}

// WordsNumber returns the total number of words in the quote.
func WordsNumber() int {
	return wordsCount
}

// PrintMap prints the whole map in a readable way
func PrintMap() {
	fmt.Println("There are", wordsCount, "words in total.")
	for k, v := range Wordsmap {
		fmt.Println(k, v)
	}
}
