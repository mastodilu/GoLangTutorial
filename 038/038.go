package main

import (
	"fmt"
)

func foo() string {
	return "hi"
}

func main() {
	a := func() func() int {
		return func() int {
			return 42
		}
	}
	fmt.Printf("type of a: %T\n", a)

	b := a()
	fmt.Printf("Type of b: %T\n", b)

	c := b()
	fmt.Printf("Type of c: %T\n", c)

	// -----------------------------------

	f1 := func() int {
		return func() int {
			return 5
		}()
	}
	fmt.Println(f1())

	// -----------------------------------

	f2 := func() func() int {
		return func() int {
			return 5
		}
	}
	fmt.Println(f2()())
}
