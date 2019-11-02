# Vari formati numerici e non

La funzione `fmt.Printf()` puo' essere usata per stampare una stessa variabile in formati diversi.

```Go
func main() {
    x := 10
    fmt.Printf("%T\n", x)  // tipo di variabile
    fmt.Printf("%b\n", x)  // binario
    fmt.Printf("%x\n", x)  // esadecimale
    fmt.Printf("%#x\n", x) // 0xEsadecimale
    fmt.Printf("%v\n", x)  // valore normale

    /*
        %v	the value in a default format
        when printing structs, the plus flag (%+v) adds field names
    */
}
```

`%v` rappresenta il formato standard di una variabile. Un numero viene stampato come numero decimale, una stringa come stringa, una struct come sequenza di valori.

Il formato `%+v` pero' stampa le struct aggiungendo i nomi dei campi prima dei valori, ad esempio:

```Go
data := struct {
    Name, Surname string
    Age           int
    Sports        []string
}{
    Name:    "Giovanno",
    Surname: "Antotognio",
    Age:     -42,
    // slice (array variabile) di stringhe
    Sports:  []string{"calcio", "nuoto", "mangiatore di hotdog"},
}

fmt.Println(data)
fmt.Printf("%v \n", data)
fmt.Printf("%+v \n", data)
```

Output

```cmd
{Giovanno Antotognio -42 [calcio nuoto mangiatore di hotdog]}
{Giovanno Antotognio -42 [calcio nuoto mangiatore di hotdog]} 
{Name:Giovanno Surname:Antotognio Age:-42 Sports:[calcio nuoto mangiatore di hotdog]}
```
