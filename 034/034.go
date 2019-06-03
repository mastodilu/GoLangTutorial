package main

import (
	"fmt"
)

// 💥 interfacce
type human interface {
	walk(int)
}

type person struct {
	first, last string
	age         int
}

//tipo person implementa interfaccia human
func (p person) walk(steps int) {
	fmt.Println("I did", steps, "steps")
}

func doStuff(h human, steps int) {
	switch h.(type) {
	case person:
		// 💥 assertion
		fmt.Println(h.(person).first, h.(person).last) // passo dall'interfaccia al tipo che implementa l'interfaccia
		h.walk(steps)
	default:
		fmt.Println("Default case...")
	}
}

func main() {
	matteo := person{
		first: "matteo",
		last:  "dilu",
		age:   25,
	}
	doStuff(matteo, 5)
}
