package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels")
	channel := make(chan string)

	// i canali bloccano l'esecuzione finchè sender e receiver non possono eseguirsi contemporaneamente
	go func() {
		channel <- "ciao"
	}()

	fmt.Println(<-channel)
}
