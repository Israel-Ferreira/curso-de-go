package main

import (
	"fmt"
)

const englishHelloPrefixer = "Hello, "

// Hello  returns "Hello, ${name}"
func Hello(name string) string {
	if name == "" {
		name = "world!"
	}

	return englishHelloPrefixer + name
}

func main() {
	fmt.Println(Hello("Israel"))
}
