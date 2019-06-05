package person

import (
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	Icecreams []string
}

func (p Person) String() string {
	var s string
	s = fmt.Sprint(p.Firstname, " ", p.Lastname, " ", "likes ")
	for _, icecream := range p.Icecreams {
		s += icecream + " "
	}
	return s
}
