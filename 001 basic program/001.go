package main // nome del package

/*
	- il nome del file non influisce sul nome del progetto o del package
	- ogni progetto deve avere UNO E UNO SOLO package main
	- ogni progetto deve avere un solo func main()
	- main() deve stare nel package main
*/

// package importati e usati nel progetto
import (
	"fmt"
)

// funzione main, viene eseguita per prima avviando il progetto con 'go run nome_file.go'
func main() {
	fmt.Println("Hey")
}
