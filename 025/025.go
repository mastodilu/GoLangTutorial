package main

import (
	"fmt"
)

func main(){
	// crea una mappa
	peopleLikings := map[string][]string{
		"matteo": {"gelati", "il cibo in generale", "musica", "netflix"},
		"alberto": {"calcio", "il suo lavoro", "essere il numero 1"},
	}
	fmt.Println(peopleLikings)

	// aggiunge un record
	peopleLikings["giovanno"] = []string{"peperonata", "la palestra"}

	// stampa in modo leggibile
	for k, v := range peopleLikings {
		fmt.Println(k)
		for j, w := range v {
			fmt.Println("\t", j, w)
		}
	}
}
