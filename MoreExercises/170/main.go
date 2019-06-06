package main

import (
	"fmt"
	"sync"
)

func reader(channel <-chan int, numbers *[]int) {
	defer wg.Done()

	for msg := range channel {
		*numbers = append(*numbers, msg)
	}
}

func writer(channel chan<- int) {
	defer wg.Done()
	defer close(channel)

	for i := 0; i < 100; i++ {
		channel <- i
	}
}

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	var numbers []int

	wg.Add(2)
	go reader(c, &numbers)
	go writer(c)
	wg.Wait()

	for _, x := range numbers {
		fmt.Print(x, "-")
	}
	fmt.Println()

	fmt.Println("Bye")
}
