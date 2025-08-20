package main

import (
	"fmt"
	"runtime"
	"sync"
)

// before waitgroup , code was unreliable, had to guess how long the goroutines would take to finish
// waiigroup ensures that the main function waits for the exact moments wg.Done() has been called by foo ()
// basically ensures foo which has isnt own goroutine gets fully printed out

var wg sync.WaitGroup // using a synchronisation primitive to wait for a collection of goroutines to finish

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // decreases the counter by 1
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func intro() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(1) // we're waiting for one thing
	go foo()  // func is launching into its own goroutine. -> code is now concurrent ( have now concurrent design pattern because of the keyword go)
	// go is non-blocking and so main goroutine does not wait for foo to start or finish. Once main finishes, program terminates regardless of other goroutines
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // once counter is 0 , it will no longer block the main function

}

// concurrent is about tasks happening at the same time rather than sequential
// paraellism is when both tasks run at the same time at the compiler ( but that is only possible with multiple CPUs)
// you can
