package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

// won't marshal if fields are in lowercase
// struct doesn't even have to be in capital ( minimum are the fields )
type person struct {
	First string
	Last  string
	Age   int
}

// struct for JSON to unmarshal into
type User struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("JSONNN")

	// marshal turns struct into JSON format

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Mis",
		Last:  "Moneypenny",
		Age:   32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// unmarshal convers JSON to struct

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Mis","Last":"Moneypenny","Age":32}]`
	bx := []byte(s)
	fmt.Println()

	var user []User

	err = json.Unmarshal(bx, &user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data from JSON:", user)

	// for funsies

	for i, v := range user {
		fmt.Println("\nPERSON NUMBER :", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	// how to sort values using the sort package
	p3 := person{
		First: "Q",
		Last:  "Ham",
		Age:   64,
	}

	p4 := person{
		First: "M",
		Last:  "Pine",
		Age:   11,
	}

	fmt.Println("-------------------------")

	people = []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	// bcrypt = encrypts password into something but output will always be the same

	pw := `password123`
	bsx, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	pw2 := `password123`
	err = bcrypt.CompareHashAndPassword(bsx, []byte(pw2))
	if err != nil {
		fmt.Println("Failed to login sis")
		return
	}
	fmt.Println("Sucessfully logged in")

}
