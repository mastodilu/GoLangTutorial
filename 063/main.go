package main

import (
	"fmt"

	"./rolldice"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Rolled the dice:", rolldice.RollDice())
	}
}
