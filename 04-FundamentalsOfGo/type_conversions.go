package main

import "fmt"

func Type_conversions() {
	fmt.Println("This is quick note on conversions in Go 🌍")

	var Yasmin int = 24

	fmt.Println(Yasmin)

	z := float32(Yasmin)
	fmt.Println(z)

	// I cant convert from int to float but I can't do int to string ?
	// looks like there are certain conversions I can do and some I can't

}
