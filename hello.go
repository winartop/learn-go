package main

import (
	"fmt"
)

/* const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

-------

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
} */

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	/* 	if language == spanish {
	   		return spanishHelloPrefix + name
	   	}

	   	if language == french {
	   		return frenchHelloPrefix + name
	   	}

	   	return englishHelloPrefix + name */

	return gretingPrefix(language) + name
}

func gretingPrefix(language string) (prefix string) {
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

func HelloWithSwitch(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))

	fmt.Println("---------------------------")

	fmt.Println(HelloWithSwitch("world", french))
}
