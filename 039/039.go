package main

import (
	"fmt"
)

// closure
// NB nel main incrementor viene assegnato ad una variabile che mantiene lo stato di x
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	// 💥 CLOSURE
	// A closure is a special type of anonymous function that references
	// variables declared outside of the function itself.

	//sono un buon modo per avere variabili instanziate e memorizzata da qualche parte in un blocco di codice il cuo scope è contenuto
	// è una possibile alternativa a quello che in java sarebbe:
	// un oggetto con una variabile privata al quale si accede con setter e getter

	for i := 0; i < 3; i++ {
		fmt.Print(incrementor()(), " ")
	}

	fmt.Println()

	a := incrementor()
	for i := 0; i < 3; i++ {
		fmt.Print(" a:", a())
	}

	fmt.Println()

	b := incrementor()
	for i := 0; i < 3; i++ {
		fmt.Print(" b:", b())
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Print(" a:", a())
	}

	a = incrementor()
	for i := 0; i < 3; i++ {
		fmt.Print(" a:", a())
	}

}
