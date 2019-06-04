package main

import (
	"fmt"
	"sync"
)

func writer(channel chan<- string) {
	defer wg.Done()
	defer close(channel)

	for i := 1; i <= 100; i++ {
		channel <- fmt.Sprint("msg", i)
	}
}

func reader(channel <-chan string) {
	defer wg.Done()

	canLoop := true
	for canLoop {
		msg, ok := <-channel
		if ok {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel closed. Terminating.")
			canLoop = false
		}
	}
}

var wg sync.WaitGroup

func main() {
	channel := make(chan string)

	wg.Add(2)
	go writer(channel)
	go reader(channel)
	wg.Wait()

	fmt.Println("Bye")
}
