package main

import (
	"fmt"
)

func main() {
	x := []int{1, 5, 7, 3, 8}
	fmt.Println("len", len(x), x)

	// append values
	x = append(x, -1, -5, 16)
	fmt.Println("len", len(x), x)

	// append slice
	y := []int{22, 33}
	y = append(y, x...)
	fmt.Println(y)

	//slice delete
	//cancella il terzo elemento
	y = append(y[:2], y[4:]...)
	fmt.Printf("Cancello elementi 3 e 4:\n%v", y)
}
