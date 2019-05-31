package main

import "fmt"

func main() {

	// nuovo array di int di 5 elementi inizializzato a zero
	var xx [5]int

	xx[2] = 42
	fmt.Println(xx)

	fmt.Println(len(xx))

	// ❗ idiomatic go dice di usare SLICE

	// SLICE

	// x := type{values} // composite literal
	x := []int{1, 5, 7, 3, 8}
	fmt.Println(x)
}
