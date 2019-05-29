package main

import (
	"fmt"
)

func main() {
	s := "Hello."
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// una stringa può essere vista come slice di bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	/*
		in UTF8 vengono usati da 1 a 4 byte per carattere
	*/

	// stampa in UTF8
	// ogni code point è un RUNE: rune è un alias per int32 e rappresenta un carattere in UTF8
	// NB: se un testo è salvato in 'bytes' allora non è vero che ogni carattere rappresenta un code point
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U \n", s[i])
	}

	// stampa il codice ASCII
	for i, v := range s {
		fmt.Println(i, " --> ", v)
	}
}
