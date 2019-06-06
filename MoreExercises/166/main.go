package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}
