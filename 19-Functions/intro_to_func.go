package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func intro() {

	// func ( r receiver) identifier( p parameter(s) return(s)) {...}

	// no params, no returns
	foo()
	// 1 param, no returns
	bar("Yasmin")
	// 1 param, 1 return
	aloha("Yusra")
	// 2 params, 2 returns
	dogYears("Todd", 32)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// unfurling a slice
	xi := []int{2, 4, 6, 8}
	x := sum(xi...)
	fmt.Println(x)

	// defer
	defer choo() // this will go last
	foo()        // this will go first before choo()

	// methods && functions

	p1 := person{
		first: "Jasmine",
	}

	p1.speak()

	// interfaces && polymorphism
	// interfaces = set of method signaturs
	// polymorphism = ability of a value of a certain type to also be of another type

	sa1 := secretAgent{
		person: person{
			first: "James Bond",
		},

		ltk: true,
	}

	sa1.speak()

	// now that interface Human is implemented , access to saySomething
	saySomething(sa1) // secretAgent and Person are both of type human

	// so this method above is polymorphic because any value of any type that
	// implements interface human is able to use saySomething
}

func foo() {
	fmt.Println(" I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)

}

func aloha(s string) string {
	return fmt.Sprint("They call me Mrs. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "is this old in dog years : "), age
}

// variadic paramter
// turns type into slice of type
// needs to be the final incoming parameter if more than one
// variadic parameter can be empty and it won't panic

func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	sv := 0
	for _, v := range i {
		sv += v
	}

	return sv
}

func choo() {
	fmt.Println("booo")
}
