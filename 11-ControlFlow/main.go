package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs ")
}

func main() {
	// fmt.Println("Hello , finally on control flow !!")

	// x := rand.Intn(250)

	// fmt.Printf("the value of x is %v \t", x)

	// if x <= 100 {
	// 	fmt.Println("between 0 and 100")
	// } else if x > 100 && x <= 200 {
	// 	fmt.Println("between 101 and 200")
	// } else {
	// 	fmt.Println("between 201 and 250")
	// }

	// switch {
	// case x <= 100:
	// 	fmt.Println("between 0 and 100")
	// case x > 100 && x <= 200:
	// 	fmt.Println("between 101 and 200")
	// default:
	// 	fmt.Println("between 201 and 250")
	// }

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("the value of x and y are %v %v \t", x, y)

	/*

		if x < 4 && y < 4 {
			fmt.Println("x and y are both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y are both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater or equal to 4 and less than or equal to 6")
		} else if x != 5 {
			fmt.Println("y is not 5")
		} else {
			fmt.Println("None of the previous cases were met")
		}

	*/

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
