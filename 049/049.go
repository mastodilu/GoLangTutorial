package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func foo() {
	defer wg.Done()
	//don't forget closures
	var count int
	for i := 0; i < 10; i++ {
		count++
		fmt.Println("foo", count)
	}
}

func bar() {
	defer wg.Done()
	//don't forget closures
	var count int
	for i := 0; i < 10; i++ {
		count++
		fmt.Println("bar", count)
	}
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Concurrency VS parallellism.\nConcurrent code:")
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("Goroutines", runtime.NumGoroutine(), " <== <== <== <== <== <== <==")
	wg.Wait()

	time.Sleep(time.Second * 1)
}
