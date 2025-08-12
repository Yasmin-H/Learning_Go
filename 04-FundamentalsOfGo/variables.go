package main

import "fmt"

func Variables() {
	fmt.Println("This is Section 4 !")

	// use var only for zero value
	// for anything else, use short-hand declaration

	var g string
	fmt.Println(g)

	// when you create variables , you wil always going to have to use it
	// or else it will throw an error / complain - since it's code pollution

	// multiple initilisiations ( declare and assign )

	a, b, c, _, d := 1, 3, 5, 7, 9

	fmt.Println(a, b, c, d)

	// use a blank identifier if you don't want to print a certain value
	// will then not complain for not using it

}
