package main

import (
	"fmt"
)

func main() {
	hello := Hello("world")
	fmt.Println(hello)
}

func Hello(name string) string {
	return "Hello, " + name
}
