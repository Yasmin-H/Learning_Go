package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	name string
}

func (u User) WriteOut(w io.Writer) { // we pass in the Writer interface to say we want to implement this
	_, err := w.Write([]byte(u.name))
	if err != nil {
		log.Fatal()
	}
}

func writer_interface() {

	fmt.Println("--------------")

	// The Writer Interface

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	b := []byte("Hello Everyone")

	_, err = file.Write(b)
	if err != nil {
		log.Fatal()
	}

	// buffers ( = temporary storage area for bytes)

	b2 := bytes.NewBufferString("Heya")
	fmt.Printf("type is : %T\n", b2)
	fmt.Printf("type is : %T\n", b2.String())
	b2.Reset()
	b2.WriteString("Never mind")
	fmt.Println(b2)
	b2.WriteString(" HAHAH")
	fmt.Println(b2)

	// Writing to a file or buffer

	p1 := User{
		name: "Yasmin",
	}

	var b3 bytes.Buffer

	p1.WriteOut(file) // write to file
	p1.WriteOut(&b3)  // write to buffer

	fmt.Println(b3.String())
}
