package main

import (
	"fmt"

	"../a/person"
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

	peopleLikings := map[string]person.Person{
		alberto.Firstname: alberto,
		matteo.Firstname:  matteo,
	}
	// peopleLikings["alberto"] = alberto
	// peopleLikings["matteo"] = matteo

	for _, v := range peopleLikings {
		fmt.Println(v.String())
	}

}
