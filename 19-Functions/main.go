package main

import (
	"fmt"
	"math"
)

// Square represents a square with a length and width
type Square struct {
	lenght float64
	width  float64
}

// Circle represents a circle with a radius
type Circle struct {
	radius float64
}

func (s Circle) getArea() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func (sq Square) getArea() float64 {
	return sq.lenght * sq.width
}

type Shape interface {
	getArea() float64
}

func info(s Shape) float64 {
	return s.getArea()
}

func main() {
	fmt.Println("-------Exercise-------")

	// practise variadic func

	answer := favNumbers(1, 5, 687, 5, 3)
	fmt.Println(answer)

	answer1 := addAllNumbers(1, 63, 6, 7, 3, 2)
	fmt.Println(answer1)

	fmt.Println("-------Exercise-------")

	sq1 := Square{
		lenght: 4,
		width:  5,
	}

	fmt.Println(sq1)

	s1 := Circle{
		radius: 5.5,
	}

	fmt.Println(s1)

	fmt.Println(info(sq1))
	fmt.Println(info(s1))

	fmt.Println("-------Exercise-------")

	fmt.Println(Add(5, 5))

}

func favNumbers(i ...int) []int {
	var xs []int
	xs = append(xs, i...)

	return xs
}

// Function returns the sum of all integers as an int
func addAllNumbers(i ...int) int {
	in := 0
	for _, v := range i {
		in += v
	}

	return in
}

// for testing purposes

// Add prints out the sum of 2 integers
func Add(a int, b int) int {
	return a + b
}
