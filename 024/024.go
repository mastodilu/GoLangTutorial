package main

import (
	"fmt"
)

func main() {

	// nested appends
	x := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var y []int

	fmt.Println(x)
	y = append(y, append(x[:3], x[7:]...)...)
	fmt.Println(y)

	// -----------------------------------------------------------------------

	//esercizio: salva dati in un array multidimensionale
	a := []string{"aaa", "bbb", "ccc", "ddd"}
	b := []string{"111", "222", "333", "444"}
	ab := [][]string{a, b}
	fmt.Println(ab)

	// stesso esercizio, ma ciclando su ogni singolo elemento
	var z [2][]string
	// z := make([][]string, 2) // 🤷‍♂️non sono sicuro di che cosa faccia

	for _, value := range a {
		z[0] = append(z[0], value)
	}
	for _, value := range b {
		z[1] = append(z[1], value)
	}
	fmt.Println(">", z)
}
