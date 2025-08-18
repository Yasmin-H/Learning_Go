package main

import (
	"fmt"
	"log"
	"os"
)

func func_part_2() {
	// anonymous function

	func() {
		fmt.Println("Hello World")
	}()

	// anonymous function with a parameter

	func(s string) {
		fmt.Println("My name is Mrs.", s)
	}("Karen")

	// func expressions ( assigning func to a var or evaluating stuff to a single value)

	x := func() {
		fmt.Printf("Blackpink \n")
	}

	x()

	// returning a func

	z := func() func() int {
		return func() int {
			return 43
		}
	}

	zz := z() // z = func() func() int so it just returns a function not an int

	zz() // this returns the func that was returned from the func which is an int

	fmt.Println(zz()) // will return zz

	// fmt.Println(z())  // will not work

	// Callback ( passing in a func as an argument)
	aa := doMath(23, 23, add)
	fmt.Println(aa)

	bb := doMath(100, 25, substract)
	fmt.Println(bb)

	// Wrapper function - it wraps the standard os.Readfile function by adding custom error handling and formatting

	k, err := ReadFile("output.txt")
	if err != nil {
		log.Fatalf("error : %s", k)
	}
	fmt.Println(k)
}

func add(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}

func doMath(x int, y int, f func(int, int) int) int {
	return f(x, y)
}

func ReadFile(filename string) ([]byte, error) {
	qq, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error : %s", qq)
	}
	return qq, err
}
