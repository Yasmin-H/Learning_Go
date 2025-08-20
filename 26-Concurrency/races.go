package main

import (
	"fmt"
	"runtime"
	"sync"
)

func races() {

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0
	const gornts = 100
	var w sync.WaitGroup
	w.Add(gornts)

	for i := 0; i < gornts; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			w.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	w.Wait() // don't carry on or exit the problem until all have been dealt with

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}

// you'll notice that if you run a go run -race main.go that there are 2 data races
// A data race is a type of concurrency bug when 2 or more routines try to access the same var/ memory location at the same time
// and atleast one of those acccesses is a write operation
// it's a race to see which goroutien wins by access the memory last - making this behavior unreliable
// so to fix a data race , you wanna use a mutex
