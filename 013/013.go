package main

import (
	"fmt"
)

func main() {
	// esecizio 1: stampa come testo i numeri da 0 a 200

	for i := 0; i < 200; i++ {
		fmt.Printf("%v --> %c\n", i, i)
	}
}
