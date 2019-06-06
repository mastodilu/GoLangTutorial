package main

import (
	"fmt"
)

func update(p *int, n int) {
	*p = n
}

func edit(a []int) {
	a[0] = 1000
}

func main() {
	x := 42
	fmt.Println("x before:", x)
	update(&x, 666)
	fmt.Println("x after:", x)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1, 2, 3, 4, 5]
	edit(arr)
	fmt.Println(arr) // [1000, 2, 3, 4, 5]
}
