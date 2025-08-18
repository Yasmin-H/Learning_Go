package main

import "fmt"

func inDelta(i *int) { // asking for pointer to address
	*i = 50
}

// value semantics
func addOne(i int) int {
	return i + 1
}

// pointer semantics

func addOneP(i *int) {
	*i += 1
}

// exercise below

var (
	ab, b, c *string
	d        *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	ab = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	x := 42
	fmt.Println(x)  // gives value of x
	fmt.Println(&x) // gives address of x ( pointer to x / *int).
	// & = ampersand

	fmt.Println("---------------")

	y := 43
	fmt.Println(y)
	inDelta(&y) // give pointer to address
	fmt.Println(y)

	fmt.Println("---------------")

	a := 1
	fmt.Println(a)
	aa := addOne(a)
	fmt.Println(aa)
	fmt.Println(a) // value won't change as the changed value been assigned to new var aa
	bc := 1
	fmt.Println(b)
	addOneP(&bc)   // no assign to new variable as it deals with pointers
	fmt.Println(b) // original value has changed

	fmt.Println("--------Exercise--------")

	fmt.Println(*ab)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

	fmt.Println(&ab)
	fmt.Println(&b)
	fmt.Println(&c)
	fmt.Println(&d)

}
