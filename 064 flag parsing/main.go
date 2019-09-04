package main

import (
	"flag"
	"fmt"
)

// collega i flag a variabili già esistenti
var treFlag string

func init() {
	flag.StringVar(&treFlag, "tre", "_", "third flag value")
}

func main() {

	fmt.Println("\nFlags before parsing\n")

	// crea e inizializza delle variabili flag (che restituiscono puntatori)
	// sono puntatori!
	var unoFlag = flag.Int("uno", -1, "first flag value")
	var dueFlag = flag.String("due", "_", "second flag value")
	fmt.Printf("%v %v\n", "uno", *unoFlag)
	fmt.Printf("%v %v\n", "due", *dueFlag)

	// treFlag non è un puntatore!
	fmt.Printf("%v %v\n", "tre", treFlag)

	// parse della command line. Associo flag e valori parsati.
	flag.Parse()

	fmt.Printf("\nFlags after parsing\n\n")
	fmt.Printf("%v %v\n", "uno", *unoFlag)
	fmt.Printf("%v %v\n", "due", *dueFlag)
	fmt.Printf("%v %v\n", "tre", treFlag)

}
