package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	chinese            = "Chinese"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	chineseHelloPrefix = "NiHao, "
)

func main() {
	hello := Hello("", "")
	fmt.Println(hello)
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
