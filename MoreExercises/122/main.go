package main

import (
	"fmt"
)

func tryClosure() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	a := tryClosure()
	b := tryClosure()

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("a", a())
		} else {
			fmt.Println("b", b())
		}
	}
}
