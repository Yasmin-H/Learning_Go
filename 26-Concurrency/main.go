package main

import (
	"fmt"
	"runtime"
	"sync"
)

// you wanna use mutex so one goroutine at a time can lock it , access it and then when it's done , others can have access to it too
// can achieve a similar thing using the pkg Atomic ( subpackage of sync)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0
	// var counter_Atomic int64
	const gornts = 100
	var w sync.WaitGroup
	w.Add(gornts)

	var mu sync.Mutex

	for i := 0; i < gornts; i++ {
		go func() {
			mu.Lock() // locks for reading and writing
			// atomic.AddInt64(&counter_Atomic, 1)  // if doing atomic way - you're writing to say you want the counter to go up by 1
			// atomic.LoadInt64(&counter_Atomic) // this reads the counter
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			w.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	w.Wait() // don't carry on or exit the problem until all have been dealt with

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}

// output

// up to 3 goroutines happening at a time
// total count went up to 100
// since we dealt with mutex, count is more accurate
