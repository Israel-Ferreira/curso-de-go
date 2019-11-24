package main

import (
	"testing"
)


func assertCorrectMessage(t *testing.T,got,want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got,want)
	}
}


func TestHello(t *testing.T) {
	t.Run("Test with name", func(t *testing.T) {
		got := Hello("Israel","English")
		want :=  "Hello, Israel"
		assertCorrectMessage(t,got,want)
	})

	t.Run("Test without name", func(t *testing.T){
		got := Hello("","English")
		want :=  "Hello, world!"
		assertCorrectMessage(t,got,want)
	})



	t.Run("In English", func(t *testing.T){
		got := Hello("Alexander","English")
		want := "Hello, Alexander"
		assertCorrectMessage(t,got,want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Carlos","Spanish")
		want := "Hola, Carlos"
		assertCorrectMessage(t,got,want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Erik Jaquin", "French")
		want := "Bonjour, Erik Jaquin"
		assertCorrectMessage(t,got,want)
	})
}


