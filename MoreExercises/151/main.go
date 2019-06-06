package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f() {
	defer wg.Done()

	mu.Lock()
	mycount := counter // salva il contatore
	runtime.Gosched()  // cambia processo

	mycount++
	counter = mycount // sovrascrive il contatore
	mu.Unlock()
}

var counter int
var wg sync.WaitGroup
var mu sync.Mutex // risolve la race condition

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
