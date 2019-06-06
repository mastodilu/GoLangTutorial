package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels")
	channel := make(chan string, 1)
	//c'è spazio per un solo messaggio prima che il canale blocchi sender in attesa di un receiver

	// i canali bloccano l'esecuzione finchè sender e receiver non possono eseguirsi contemporaneamente
	channel <- "ciao"

	fmt.Println(<-channel)
}
