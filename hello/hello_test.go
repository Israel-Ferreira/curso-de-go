package main

import (
	"testing"
)


func TestHello(t *testing.T) {

	assertCorrectMessage :=  func(t *testing.T,got,want string){
		t.Helper()

		if got != want {
			t.Errorf("got %s, want %s", got,want)
		}
	}

	t.Run("Test with name", func(t *testing.T) {
		got := Hello("Israel")
		want :=  "Hello, Israel"
		assertCorrectMessage(t,got,want)
	})

	t.Run("Test without name", func(t *testing.T){
		got := Hello("")
		want :=  "Hello, world!"
		assertCorrectMessage(t,got,want)
	})
}


