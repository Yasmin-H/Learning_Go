package main

import "fmt"

func Binary_Hexadecimal() {
	fmt.Println("Helloo ☺️")

	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", a, a, a)
	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", b, b, b)
	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", c, c, c)
	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", d, d, d)
	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", e, e, e)
	fmt.Printf("normal value : %v \t  value in binary : %b \t value in hexadecimal:  %#X\n", f, f, f)

}
