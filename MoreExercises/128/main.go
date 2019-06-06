package main

import (
	"fmt"
)

type Person struct {
	Firstname, Lastname string
}

func changeMe(p *Person) {
	// (*p).Firstname == p.Firstname
	p.Firstname = "oettaM"
}

func main() {
	matteo := Person{"Matteo", "Dilu"}
	fmt.Println(matteo)
	changeMe(&matteo)
	fmt.Println(matteo)
}
