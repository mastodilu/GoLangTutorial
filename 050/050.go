package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Race condition.")

	counter := 0
	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			v := counter
			counter++
			fmt.Println(v)
		}()
	}

	wg.Wait()
}
