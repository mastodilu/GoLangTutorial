package main

import (
	"fmt"
)

func main() {

	// IF

	// inizializza a prima di verificare la condizione dell'if.
	// Lo scope di a è il blocco if
	if a := 3; a == 3 {
		fmt.Println(a)
	}
	// fmt.Println(a) // a non esiste

	// x := 1
	// if x == 1 {
	// 	fmt.Println(x)
	// }

	// IF ELSE con inizializzazione
	x := 1
	if x == 0 {
		fmt.Println("x = 0")
	} else if a := 1; a == 1 {
		fmt.Println("a = 1")
	}

}
