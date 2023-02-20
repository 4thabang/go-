package main

import (
	"fmt"
)

func main() {
	sum := Adder(2, 10)
	fmt.Println(sum)
}

func Adder(x, y int) int {
	return x * y
}
