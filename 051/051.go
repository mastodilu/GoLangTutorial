package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex // 💥 accesso in mutua esclusione alla zona critica

func main() {
	fmt.Println("Race condition.")

	counter := 0
	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			mu.Lock() // 💥
			v := counter
			counter++
			mu.Unlock() // 💥
			fmt.Println(v)
		}()
	}

	wg.Wait()
}
