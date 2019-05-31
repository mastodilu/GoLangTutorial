package main

import (
	"fmt"
)

func main() {
	x := []int{1, 5, 7, 3, 8}
	fmt.Println("len", len(x), x)

	fmt.Println("x[0:len(x)]\t\t", x[0:len(x)])
	fmt.Println("x[:len(x)]\t\t", x[:len(x)])
	fmt.Println("x[1:]\t\t", x[1:])
	fmt.Println("x[:]\t\t", x[:])
	fmt.Println("x[1 : len(x)-1]\t\t", x[1:len(x)-1])
}
