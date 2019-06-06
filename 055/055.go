package main

import (
	"fmt"
)

func reader(channel <-chan string) {
	for msg := range channel {
		fmt.Println(msg)
	}
	fmt.Println("Channel closed.")
}

func writer(channel chan<- string) {
	for i := 1; i <= 100; i++ {
		channel <- fmt.Sprintf("msg %d", i)
	}
	close(channel)
}

func main() {
	fmt.Println("Channel range")
	channel := make(chan string)

	go writer(channel)
	reader(channel)

	fmt.Println("Bye")
}
