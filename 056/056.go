package main

import (
	"fmt"
)

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

func receive(even, odd <-chan int, quit <-chan bool) {
receiverLoop: // 💥
	for {
		select { // 💥
		case evenNum := <-even: // 💥
			fmt.Println("Even:", evenNum) // 💥
		case oddNum := <-odd: // 💥
			fmt.Println("Odd:", oddNum) // 💥
		case mustExit := <-quit: // 💥
			fmt.Println("Exit:", mustExit) // 💥
			break receiverLoop             // 💥
		}
	}
	fmt.Println("Receiver terminating.")
}

func main() {
	fmt.Println("Channel select.")

	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan bool)

	go send(evenChannel, oddChannel, quitChannel)
	receive(evenChannel, oddChannel, quitChannel)
}
