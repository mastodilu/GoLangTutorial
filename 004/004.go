package main

import (
	"fmt"
)

type interoh int

func main() {
	x := 42
	var y interoh // creo il mio tipo che rappresenta un int
	y = 10

	fmt.Printf("%d %T\n", x, x)
	fmt.Printf("%d %T\n", y, y)

	/*
		errore: tipi diversi
		💥 GO è statico: non puoi dinamicamente cambiare il tipo di una variabile
		x = y
		y = x
	*/
}
