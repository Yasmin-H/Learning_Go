package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs ")
}

func main() {
	fmt.Println("Hello , finally on control flow !!")

	// if statements

	x := rand.Intn(250)

	fmt.Printf("the value of x is %v \t", x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x > 100 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}

	// switch statmements

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("between 101 and 200")
	default:
		fmt.Println("between 201 and 250")
	}

	// for loop 1# :
	for i := 0; i < 100; i++ {
		fmt.Printf("this is iteration number %v \n", i)
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("iteration %v \t the value of x and y are %v %v \t", i, x, y)
		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater or equal to 4 and less than or equal to 6")
		case x != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("None of the previous cases were met")
		}

	}

	x1 := rand.Intn(10)
	y1 := rand.Intn(10)

	if x1 < 4 && y1 < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x1 > 6 && y1 > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x1 >= 4 && x1 <= 6 {
		fmt.Println("x is greater or equal to 4 and less than or equal to 6")
	} else if x1 != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the previous cases were met")
	}

	for i := 0; i < 42; i++ {

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("Iteration %v \t the value of x is %v\n", i, x)
		case 1:
			fmt.Printf("Iteration %v \t the value of x is %v\n", i, x)
		case 2:
			fmt.Printf("Iteration %v \t the value of x is %v\n", i, x)
		case 3:
			fmt.Printf("Iteration %v \t the value of x is %v\n", i, x)
		case 4:
			fmt.Printf("Iteration %v \t the value of x is %v\n", i, x)
		}

	}

	z := 0

	// for loop 2#
	for z < 10 {
		fmt.Printf("value of z is %v \n", z)
		z++

		// for loop 3#

		for {
			fmt.Printf("The value of z is %v\n", z)
			if z > 10 {
				break
			}
			z++
		}

		m := map[string]int{
			"James":      42,
			"MoneyPenny": 32,
		}

		for k, v := range m {
			fmt.Printf("key : %v value : %v\n", k, v)
		}

		// use lookup to find value of key

		age := m["James"]
		fmt.Println(age)

		q := m["Q"]
		fmt.Println(q)

		// use comma ok idiom to do something similar

		if v, ok := m["Q"]; !ok {
			fmt.Println("There is no Q, and here is the zero value of int", v)
		}

		var c int = 1
		for i := 0; i < 100; i++ {
			if x := rand.Intn(5); x == 3 {
				fmt.Printf("iteration %v \t total count %v \t x is %v", i, c, x)
				c++
			}
		}

	}

}
