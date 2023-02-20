package main

import (
	"fmt"
	"strings"
)

func main() {
	iter := Iteration("A", 5)
	fmt.Println(iter)
}

func Iteration(letter string, iteration int) string {
	var stringBox string
	for i := 0; i < iteration; i++ {
		stringBox += letter
	}
	return strings.ToUpper(stringBox)
}
