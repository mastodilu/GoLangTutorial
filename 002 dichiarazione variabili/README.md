# Dichiarazione di variabili

Le variabili possono essere locali o globali, pubbliche o private.

Per dichiarare una variabile si usa la keyword `var` oppure la sintassi `:=`.

```Go
var foo float
```

dichiara la variabile foo e le assegna il valore 0 del tipo float.

```Go
foo := 0.14
```

dichiara la variabile foo di tipo float e le assegna il valore 0.14.

Altri metodi alternativi sono:

```Go
// 1
var name, surname string
// 2
name, surname := "matteo", "dilu"
// 3
var name, surname string = "matteo", "dilu"
// 4
var (
    a, b, c, d = 1, 2, 3, 4
    name = "matteo"
)
```

## Variabili locali o globali

- globali: sono dichiarate fuori da una funzione e sono visibili all'interno del package.
- locali: sono dichiarate all'interno di una funzione, e visibili soltanto nella funzione in cui vengono dichiarate.

```Go
// variabile globale
var a = 10
var b int

func main() {
    c := "ciao"
}
```

## Variabili pubbliche e private

La visibilita' nel progetto di una variabile e' determinata da come viene scritta:

- se una variabile globale viene scritta con l'iniziale maiuscola allora e' visibile in tutto il progetto e puo' essere chiamata importanto il package a cui appartiene.

```Go
var name string // private variable
var Email string // public variable
```

## Ridichiarare le variabili

E' possibile ridichiarare una variabile gia' esistente se hanno scope diverso:

```Go
var a = "global A"

func sayA() {
    fmt.Println(a)
}

func main() {
    a := "local A"
    fmt.Println(a) // print "global A"
    sayA() // print "local A"
}
```

Ma questo genera un **errore**:

```Go
// error example
func main() {
    a := "A"
    a := "another A"
    // or
    var a string
    a := "new A"
}
```
