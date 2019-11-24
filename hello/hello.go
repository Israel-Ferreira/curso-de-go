package main

import (
	"fmt"
)

const englishHelloPrefixer = "Hello, "

// Hello  returns "Hello, ${name}"
func Hello(name, language string) string {
	if name == "" {
		name = "world!"
	}


	return greetingPrefix(language) + name
}


func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}

	return
}

func main() {
	fmt.Println(Hello("Israel","English"))
}
