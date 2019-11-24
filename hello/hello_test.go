package main

import (
	"testing"
)


func TestHello(t *testing.T) {
	t.Run("Test with name", func(t *testing.T) {
		got := Hello("Israel")
		want :=  "Hello, Israel"

		if got != want {
			t.Errorf("got %s, want %s",got,want)
		}
	})

	t.Run("Test without name", func(t *testing.T){
		got := Hello("")
		want :=  "Hello, world!"

		if got != want {
			t.Errorf("got %s, want %s",got,want)
		}
	})
}


