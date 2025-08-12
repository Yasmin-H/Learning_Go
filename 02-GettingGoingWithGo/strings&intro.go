package main

import "fmt"

func Practice() {
	fmt.Println("Hello Gophers ☺️")

	const name, age = "Yasmin", 24
	const float = 3.5

	// to find the Type of a value = %T
	// for integers = %d
	// for strings = %s
	// for floats = %e
	// for defaul value == %v ( can use this for both)

	fmt.Printf("%s is %d years old \n", name, age)
	fmt.Printf("%e is a float whereas this is a %T", float, name)

	fmt.Println(`
	This is a raw string literal.
	So it basically means that it will literally print this ,
	in the exact same layout as this.
	But you have to use backticks.`)
}
