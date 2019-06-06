// Package person contiene tipi e funzioni per un ipotetico tipo Person.
package person

import (
	"fmt"
)

// Person è un tipo che rappresenta una persona
type Person struct {
	Firstname string
	Lastname  string
}

// Speak stampa Lastname e Firstname del tipo Person passato
func (p *Person) Speak() string {
	return fmt.Sprintf("I'm %s %s", (*p).Lastname, (*p).Firstname)
}
