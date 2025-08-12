package main

import (
	"fmt"

	"github.com/Yasmin-H/puppy"
)

func GoModules() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1, s2)

	// also like this

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
	fmt.Println("Hello we're working on versions here too ")
	fmt.Println(puppy.From11())

}
