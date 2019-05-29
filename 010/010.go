package main

import (
	"fmt"
)

const (
	_  = iota             // 0
	kb = 1 << (iota * 10) // 10
	mb = 1 << (iota * 10) // 20
	gb = 1 << (iota * 10) // 30
	// kb = 1024
	// mb = 1024 * kb
	// gb = 1024 * mb
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	// BIT SHIFTING

	// shifta il numero binario x a sinistra di una posizione (aggiunge uno 0 in coda: moltiplica per 2)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	// shifta il numero binario a destra di una posizione (rimuove uno 0 in coda: divide per 2)
	z := y >> 1
	fmt.Printf("%d\t\t%b\n", z, z)

	fmt.Printf("%d -> %b\n", kb, kb)
	fmt.Printf("%d -> %b\n", mb, mb)
	fmt.Printf("%d -> %b\n", gb, gb)
}
