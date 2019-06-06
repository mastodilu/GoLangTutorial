package main

import (
	"fmt"

	"./person"
)

func main() {
	alberto, matteo := person.Person{
		Firstname: "alberto",
		Lastname:  "dilu",
		Icecreams: []string{"fragola", "cioccolato", "fango"},
	}, person.Person{
		Firstname: "matteo",
		Lastname:  "dilu",
		Icecreams: []string{"yogurt", "caffè", "pesca"},
	}

	fmt.Println(alberto.String())
	fmt.Println(matteo.String())
}
