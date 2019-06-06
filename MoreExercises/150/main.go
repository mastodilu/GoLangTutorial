package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f() {
	defer wg.Done()

	mycount := counter // salva il contatore
	runtime.Gosched()  // cambia processo

	mycount++
	counter = mycount // sovrascrive il contatore
}

/*

			go run -race main.go

==================
WARNING: DATA RACE
Read at 0x0000006062f8 by goroutine 7:
  main.f()
      C:/Users/matteo.dilucchio/Documents/Progetti random/GoLangTutorial/MoreExercises/150/main.go:12 +0x6d

Previous write at 0x0000006062f8 by goroutine 6:
  main.f()
      C:/Users/matteo.dilucchio/Documents/Progetti random/GoLangTutorial/MoreExercises/150/main.go:16 +0x8e

Goroutine 7 (running) created at:
  main.main()
      C:/Users/matteo.dilucchio/Documents/Progetti random/GoLangTutorial/MoreExercises/150/main.go:28 +0x6f

Goroutine 6 (finished) created at:
  main.main()
      C:/Users/matteo.dilucchio/Documents/Progetti random/GoLangTutorial/MoreExercises/150/main.go:28 +0x6f
==================
370
Found 1 data race(s)
exit status 66
*/

var counter int
var wg sync.WaitGroup

func main() {
	const max int = 1000

	// race condition
	wg.Add(max)
	for i := 0; i < max; i++ {
		go f()
	}
	wg.Wait()
	// -----------

	fmt.Println(counter)
}
