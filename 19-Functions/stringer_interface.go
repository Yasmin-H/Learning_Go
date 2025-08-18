package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title fo the book is", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM FUTURE", s.String())
} // adds additional message before the String function is called on any type

func stringer() {

	// Exploring the stringer interface

	b := book{
		title: "ACOTAR",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

	// Wrapper func for logging ( wrapper func is when you want to add something extra to a known func)
	// either by wrapping it around with another func or message
	// useful when logging, handling errors and timing

	log.Println(b) // what if you want to add something additional to the log function
	logInfo(b)     // now log has additional message we want

}
