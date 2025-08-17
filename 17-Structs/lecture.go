package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func lecture() {
	// difference between maps and structs are that structs can hold
	// values of different types
	// and you use structs to create your own type

	p1 := person{
		first: "Yasmin",
		last:  "Haidar",
		age:   24,
	}

	p2 := person{
		first: "Yusra",
		last:  "Haidar",
		age:   23,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.age)

	// embedded structs

	sa1 := secretAgent{
		person: person{
			first: "Dodo",
			last:  "Haidar",
			age:   18,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.person.first)

	// anonymous structs - note how a struct was made and then using a literal the values were defined
	// the struct is not named hence anonymous

	r1 := struct {
		name     string
		sentence string
	}{
		name:     "So Random",
		sentence: "OG Disney Show",
	}

	fmt.Println(r1)

	// composition
	// = way of structuring and building complex types by combining multiple simpler types
	// provides a way to create relationships between different types without the need for traditional inheritance
}
