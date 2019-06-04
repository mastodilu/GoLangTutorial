package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hey there.")
	fmt.Fprintln(os.Stdout, "hey there.")
	io.WriteString(os.Stdout, "hey there.\n")
	os.Stdout.Write([]byte("Hey there."))
}
