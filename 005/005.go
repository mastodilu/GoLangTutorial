package main

import (
	"fmt"
)

type interoh int

func main() {
	x := -99
	var y interoh // creo il mio tipo che rappresenta un int
	y = 10

	fmt.Printf("%d %T\n", x, x)
	fmt.Printf("%d %T\n", y, y)

	// per assegnare il tipo "interoh" al tipo int posso effettuare una conversione
	// x = int(y)
	// fmt.Printf("%d %T\n", x, x)

	// per assegnare il tipo int al tipo "interoh" posso effettuare una conversione
	// 💥 il metodo interoh(var) è implicito
	y = interoh(x)
	fmt.Printf("%d %T\n", y, y)
}
