package main

import "fmt"

func main() {

	x := []int{1, 5, 7, 3, 8}
	fmt.Println("len", len(x), x)

	fmt.Println(x[0], x[1])

	for i := 0; i < len(x); i++ {
		fmt.Println(x)
	}

	for i := range x {
		fmt.Println("index:", i)
	}

	for i, value := range x {
		fmt.Println(i, ":", value)
	}

	for _, value := range x {
		fmt.Println(value)
	}
}
