package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Race condition.")

	var counter int64 // 💥 per il package atomic

	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)            // 💥 write
			fmt.Println(atomic.LoadInt64(&counter)) // 💥 read
		}()
	}

	wg.Wait()
}
