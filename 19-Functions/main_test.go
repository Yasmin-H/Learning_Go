package main

import (
	"log"
	"testing"
)

// unit testing

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubstract(t *testing.T) {
	got := substract(7, 5)
	want := 2
	if got != want {
		log.Fatalf("error - want %d and got %d", want, got)
	}
}
