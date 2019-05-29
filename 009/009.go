package main

import (
	"fmt"
)

/*
	iota è una keyword riservata dispobibile all'interno di "const parentetizzato"
	che si incrementa di 1 rispetto al precedente iota, e comincia da 0.
*/

const (
	a = iota
	b = iota
	c = iota
	d = iota
)

const (
	e = iota
	f = iota
	g = iota
)

const (
	nb1 = iota + 1
	nb2 = iota
)

func main() {

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// -----------------
	fmt.Println()

	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)

	// -----------------
	fmt.Println()

	fmt.Printf("%v\n", nb1)
	fmt.Printf("%v\n", nb2)
}
