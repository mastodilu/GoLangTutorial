package main

import (
	"fmt"
)

func main() {
	c, quit := gen() // genera i numeri e invia al canale

	receive(c, quit)

	fmt.Println("about to exit")
}

func receive(c <-chan int, quit <-chan bool) {

	var msg int
loop:
	for {
		select {
		case msg = <-c:
			fmt.Println(msg)
		case <-quit:
			break loop
		}
	}
}

func gen() (<-chan int, <-chan bool) {
	c := make(chan int)
	q := make(chan bool)

	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
			fmt.Println(i)
		}
		q <- true
	}()

	return c, q
}
