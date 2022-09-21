package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func main() {
	hello := Hello("")
	fmt.Println(hello)
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
