package main

import (
	"fmt"
)

func main() {
	mydata := struct {
		words []string
		pairs map[string][]string
	}{
		words: []string{"AAA", "BBB", "CCC"},
		pairs: map[string][]string{
			"AAA": []string{"one", "two", "three"},
			"BBB": []string{"four", "five", "six"},
			"CCC": []string{"seven", "eight", "nine"},
		},
	}

	for _, id := range mydata.words {
		fmt.Println(id)
		for _, v := range mydata.pairs[id] {
			fmt.Printf("\t%v\n", v)
		}
	}
}
