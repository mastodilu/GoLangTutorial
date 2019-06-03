package main

import (
	"fmt"
)

func foo(n int) {
	fmt.Println("foo", n)
}

func bar(n int) {
	fmt.Println("bar", n)
}

func main() {
	// defer() ritarda l'esecuzione di una funzione al termine della funzione in cui viene chiamata
	defer foo(1)
	defer bar(2)
	bar(3)
	foo(4)
}
