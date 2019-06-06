package main

import (
	"fmt"
)

func inception() func() string {
	f := func() string {
		return "Inception!"
	}
	return f
}

func main() {
	f := inception()
	fmt.Printf("inception type: %T, value: %v\n", f, f())
}
