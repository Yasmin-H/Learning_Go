package main

import "fmt"

func main() {

	fmt.Println("------Exercise 1------")

	z := [5]int{}

	for i := 0; i < 5; i++ {
		z[i] = i
	}
	for i, v := range z {
		fmt.Printf("index position : %v \t value is : %v \t and its type is %T\n", i, v, v)
	}

	fmt.Println("------Exercise 2------")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range x {
		fmt.Printf("the value is : %v and the type is %T\n", v, v)
	}

	fmt.Println("------Exercise 3------")

	slice1 := x[:5]
	fmt.Println(slice1)

	slice2 := x[5:]
	fmt.Println(slice2)

	slice3 := x[2:7]
	fmt.Println(slice3)

	slice4 := x[1:6]
	fmt.Println(slice4)

	fmt.Println("------Exercise 4------")

	append1 := append(x, 52)
	fmt.Println(append1)

	append2 := append(x, 53, 54, 55)
	fmt.Println(append2)

	y := []int{56, 57, 58, 59, 60}

	append3 := append(x, y...)
	fmt.Println(append3)

	fmt.Println("------Exercise 5------")

	// use append 3 for this exercise
	delete := append(append3[:3], append3[6:10]...)
	fmt.Println(delete)

	fmt.Println("------Exercise 6------")

	fmt.Println("------Exercise 7------")

}
