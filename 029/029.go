package main

import (
	"fmt"
)

//functions

// no parameters - no return
func foo() {
	fmt.Println("I'm in foo")
}

// one parameter - no return
func bar(msg string) {
	fmt.Println("I'm in bar:", msg)
}

// no parameter - one return
func foobar(msg string) string {
	return "you said: " + msg
}

// due parametri con lo stesso tipo
func sum(n1, n2 int) int {
	return n1 + n2
}

// restituisce più parametri
func swap(n1, n2 int) (int, int) {
	return n2, n1
}

func main() {
	foo()                     // no arguments - no return
	bar("HEY!!")              // one argument - no return
	fmt.Println(foobar("yo")) // one argument - one return

	n1, n2 := 10, 5
	fmt.Printf("%d + %d = %d\n", n1, n2, sum(n1, n2))

	fmt.Printf("swap n1:%d, n2:%d", n1, n2)
	n1, n2 = swap(n1, n2)
	fmt.Printf(" --> n1:%d, n2:%d\n", n1, n2)
}
