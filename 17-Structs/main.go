package main

import "fmt"

type user struct {
	first_name         string
	last_name          string
	fav_ice_cream_flav []string
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	color  string
}

func main() {
	fmt.Println("------- Exercise 1 -------")

	u1 := user{
		first_name:         "yusuf",
		last_name:          "Bently",
		fav_ice_cream_flav: []string{"chocolate", "mango", "passionFruit"},
	}

	u2 := user{
		first_name:         "Harris",
		last_name:          "J",
		fav_ice_cream_flav: []string{"strawberry", "pistachio", "avocado", "vanilla"},
	}

	fmt.Println(u1)
	fmt.Println(u2)

	for _, v := range u1.fav_ice_cream_flav {
		fmt.Println(v)
	}

	fmt.Println("------- Exercise 2 -------")

	m1 := map[string]user{
		u1.last_name: u1,
		u2.last_name: u2,
	}

	for _, v := range m1 {
		fmt.Println(v)
		for _, v2 := range v.fav_ice_cream_flav {
			fmt.Println(v2)
		}
	}

	fmt.Println("------- Exercise 3 -------")

	e1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Japan",
		model: "M2",
		color: "Navy",
	}

	e2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "German",
		model: "G2",
		color: "Red",
	}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e1.color)
	fmt.Println(e2.engine)

	fmt.Println("------- Exercise 4 -------")

	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Linda",
		friends: map[string]int{
			"Amy":    23,
			"Shelly": 24,
			"Zeze":   26,
			"Zack":   26,
		},
		favDrinks: []string{"Fanta", "Coke", "Shani", "Pepsi"},
	}

	fmt.Println(s1)

}
