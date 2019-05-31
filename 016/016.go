package main

import (
	"fmt"
)

func main() {
	// esecizio 1: stampa le unicode runes dell'alfabeto per 3 volte
	for i := 65; i < 91; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d --> %#U\n", i, i)
		}
		fmt.Println()
	}

	// esercizio 2: stampa il resto della divisione per 4 di tutti i numeri da 10 a 40
	for i := 0; i < 40; i++ {
		fmt.Print(i%4, " ")
	}
	fmt.Println()

	// esercizio 3: fai uno switch sulla variabile string favSport
	switch favSport := "calcio"; favSport {
	case "nuoto":
		fmt.Println("Nuoto")
	case "calcio":
		fmt.Println("Calcio")
	default:
		fmt.Println("Non trovo il tuo sport preferito")
	}
}
