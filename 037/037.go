package main

import (
	"fmt"
)

func foo() string {
	return "hi"
}

func main() {
	// 💥 questa funzione restituisce una funzione che restituisce un intero
	// il tipo restituito è: func() int
	a := func() func() int {
		return func() int {
			return 42
		}
	}
	// a ==> func() func() int
	fmt.Printf("type of a: %T\n", a)

	// valuto / eseguo la funzione a
	b := a() // b ==> func() int

	fmt.Printf("Type of b: %T\n", b)

	// valuto la funzione b
	c := b() // c ==> int

	fmt.Printf("Type of c: %T\n", c)

	//questa funzione restituisce un intero
	f1 := func() int {
		// 💥 restituisce func() int dopo averla eseguita
		return func() int {
			return 5
		}() // <-- le parentesi indicano che la funzione viene eseguita
	}
	fmt.Println(f1())

	f2 := func() func() int {
		return func() int {
			return 5
		}
	}

	fmt.Println(f2()()) // f2 restituisce func() int che viene poi eseguita
}
