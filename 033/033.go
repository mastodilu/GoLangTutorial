package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	licenseToKill bool
	person
}

func (s secretAgent) speak() {
	fmt.Println("I'm a secret agent")
}

func (p person) speak() {
	fmt.Println("I'm just a person")
}

func main() {
	matteo := person{
		first: "matteo",
		last:  "dilu",
		age:   25,
	}
	jamesBond := secretAgent{
		licenseToKill: true,
		person: person{
			first: "james",
			last:  "Bond",
			age:   32,
		},
	}

	// 💥 METODI
	jamesBond.speak()
	matteo.speak()
}
