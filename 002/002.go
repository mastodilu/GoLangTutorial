package main

import (
	"fmt"
)

// variabile globale
var a = 10
var b int

func main() {

	c := 20 // variabile locale

	// il tipo viene assegnato in fase di compilazione, non serve specificarlo
	// una variabile non inizializzata (b) contiene il valore zero di default per il suo tipo

	fmt.Println(a + c)

	//variabile ridichiarata perchè è prioritaria da quella dichiarata fuori dal main
	b := "hey"
	fmt.Println(b)

	foo()
}

func foo() {
	// b è quella dichiarata fuori dal main (variabile globale)
	// e contiene il suo valore "zero" perchè non è stata inizializzata
	fmt.Println(b)
}
