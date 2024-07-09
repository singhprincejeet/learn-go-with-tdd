package main

import (
	"fmt"
)

const (
	spanishHelloPrefix = "Hola"
	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"

	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := getGreetingPrefix(language)

	return prefix + ", " + name + "!"
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
