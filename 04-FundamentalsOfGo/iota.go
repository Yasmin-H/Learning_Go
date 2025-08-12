package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // basically 1 left shifted 10 times which is 2^10 which is equal to 1024
	MB                             // same expression byt this time by 20 times and so on
	GB
	TB
	PB
	EB
)

func Iota() {
	fmt.Println(`Hello everyone, going to learn about iota's here !`) // note that iota is used as an incrementing counter on constants ONLY
	// also very useful when bit shifting using << or >>

	fmt.Printf("decimal: %d \t\t\t binary: %b \n ", KB, KB)
	fmt.Printf("decimal: %d \t\t binary: %b \n ", MB, MB)
	fmt.Printf("decimal: %d \t\t binary: %b \n ", GB, GB)
	fmt.Printf("decimal: %d \t binary: %b \n ", TB, TB)
	fmt.Printf("decimal: %d \t binary: %b \n ", PB, PB)
	fmt.Printf("decimal: %d \t binary: %b \n ", EB, EB)

}
