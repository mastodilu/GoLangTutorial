package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func reader(c <-chan int) {
	defer wg.Done()

	canRead := true
	for canRead {
		msg, ok := <-c
		if ok {
			fmt.Println(msg)
		} else {
			canRead = false
		}
	}
}

func writer(c chan<- int) {
	defer wg.Done()
	defer close(c)

	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Millisecond)
		c <- i
	}
}

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	wg.Add(2)
	go writer(c)
	go reader(c)
	wg.Wait()

	fmt.Println("Bye")
}
