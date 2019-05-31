package main

import "fmt"

func main() {
	// make a slice: make(slice type, length, capacity)
	x := make([]int, 2, 10)
	fmt.Printf("len:%d cap:%d - %v\n", len(x), cap(x), x)

	// ❗ NB: quando esegui questo ciclo si nota come l'array riempito viene reistanziato con capacità doppia

	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("len:%d cap:%d\t%v\n", len(x), cap(x), x)
	}

}
