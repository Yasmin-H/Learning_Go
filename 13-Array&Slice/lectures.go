package main

import "fmt"

func lectures() {

	// general convention is to use slice rather than arrays are slice are build upon them and are dynamic ( grow & shrink )

	// array literal
	a := [3]int{1, 2, 3}
	fmt.Println(a) // this is type [3] int

	b := [...]string{"Hello", "Gophers"}
	fmt.Println(b) // this is type [2] string

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c) // this is type [2] int

	// a = b  // won't work because even though both int , it's different. Type is also with number

	// find number of elements in array

	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))

	fmt.Println("---------- Exercise ---------")

	dessert := [...]string{
		"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}

	fmt.Printf("there are %v elements in this array literal and the type is %T \n", len(dessert), dessert)

	// slice literal

	xs := []string{"hello", "world"}
	fmt.Println(xs)

	fmt.Println("--------- Exercise 2 ----------")

	xs2 := []string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}

	fmt.Println(xs2)
	fmt.Println(len(xs2))

	for i, v := range xs2 {
		fmt.Printf("index: %v flavour: %v\n", i, v)
	}

	// accessing values by index position

	fmt.Println("--------------------")
	fmt.Println(xs2[0])
	fmt.Println(xs2[1])
	fmt.Println(xs2[5])
	fmt.Println(xs2[10])

	// use for loop to print values

	for i := 0; i < 31; i++ {
		fmt.Println(xs2[i])
	}

	fmt.Println("-----------------")

	// Append a slice
	xi := []int{1, 2, 3}
	fmt.Println(xi)
	xi = append(xi, 4, 5, 6, 7)
	fmt.Println(xi)

	// Slicing a slice ( cut part of the slice )

	fmt.Println("-----------------")

	xp := []int{2, 4, 6, 8, 10}
	fmt.Println(xp)

	//  [inclusive:exlcusive]
	fmt.Println(xp[1:3])

	// [:exclusive]
	fmt.Println(xp[:3])

	// [inclusive:]
	fmt.Println(xp[2:])

	fmt.Println(xp[:])

	fmt.Println("---------------")

	// Deleteing a slice ( by combing append and slice)

	xx := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	xx = append(xx[:4], xx[4:]...) // function says can only take elements, by adding variadic (....) , you're unfurling the slice

	fmt.Println(xx)

	fmt.Println("--------")

	// Create a slice using make()
	xy := make([]int, 0, 10) // telling to prepare to hold 10 but not putting anything yet. Can resize later
	fmt.Println(xy)
	fmt.Println(len(xy)) // will be 0
	fmt.Println(cap(xy)) // will be 10 - can change if more than 10 added into array

}
