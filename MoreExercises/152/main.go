package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func f() {
	defer wg.Done()
	var mycount int64
	mycount = atomic.LoadInt64(&counter) // salva il contatore con un'operazione atomica
	runtime.Gosched()                    // cambia processo
	time.Sleep(time.Millisecond * 20)

	mycount++
	atomic.AddInt64(&counter, 1) // sovrascrive il contatore con un'operazione atomica
}

var counter int64
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

	fmt.Println(atomic.LoadInt64(&counter))
}
