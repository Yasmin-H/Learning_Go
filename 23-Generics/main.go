package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// both functions are essentially the same. Only difference is that one takes in an int and the other a float64

func addI(a int, b int) int {
	return a + b
}

// this is a type constriaint - this generic func can only take in these types
func addF(a, b float64) float64 {
	return a + b
}

// this is a generic type set interface
// ~ means as well their underlying type
type myNumbers interface {
	~int | ~float64
}

// can use the constrains pkg that has it done for us
type theirNumbers interface {
	constraints.Integer | constraints.Float
}

// so let's make a generic function that can take an int or float64
// genertics alllows you to write code with a set of types instead of just a single type

func addT[T int | float64](a, b T) T {
	return a + b
}

// alternatively could use a type set to put in your name

func addTQ[T myNumbers](a, b T) T {
	return a + b
}

func addX[T theirNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	fmt.Println(addI(2, 2))
	fmt.Println(addF(4, 4))

	fmt.Println("--------")

	fmt.Println(addT(2, 2))
	fmt.Println(addT(4, 4))
	fmt.Println(addTQ(4, 5.5))

	// so now can use variables with underlying type of int or float as well in genertic func ( ~)
	var n myAlias = 42
	fmt.Println(addTQ(n, 8))
	fmt.Println(addX(1, 1))

}
