package main

import (
	"fmt"
)

func main() {
	// esercizio 1: stampa in decimale, binario ed esadecimale
	const num = 17
	fmt.Println(num)                    // decimale
	fmt.Printf("%d is %b\n", num, num)  // binario
	fmt.Printf("%d is %#X\n", num, num) // esadecimale

	// esercizio 2: crea due costanti, una con e una senza tipo
	const con1 = "ciao"
	const con2 int = 42

	// esercizio 3:
	// - crea una variabile int
	// - stampa in decimale, binario ed esadecimale
	// - esegui dei bit shifting
	// - stampa in decimale, binario ed esadecimale
	n := 98
	fmt.Printf("%d - %#X - %b\n", n, n, n)
	x := n << 1 // moltiplica per 2
	fmt.Printf("%d - %#X - %b\n", x, x, x)
	x = x << 1 // moltiplica per 2
	fmt.Printf("%d - %#X - %b\n", x, x, x)
	m := 100
	fmt.Println("m is", m)
	fmt.Printf("(moltiplica per 4) m << 2 : %d\n", m<<2)
	fmt.Printf("(divisione intera per 32) m >> 5 : %d\n", m>>5)

	// esercizio 4: crea un raw string literal
	s := `ciao
bau`
	fmt.Println(s)

	// esercizio 5: usa iota e stampa 4 anni consecutivi
	const (
		// baseYear = 2019
		year1 = 2019 + iota
		year2 = 2019 + iota
		year3 = 2019 + iota
		year4 = 2019 + iota
	)
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
