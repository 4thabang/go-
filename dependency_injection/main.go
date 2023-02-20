package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	name := Name("Thabang")
	name.Greet(os.Stdout)
}

type Name string

func (n Name) Greet(writer io.Writer) {
	fmt.Fprintf(writer, "Hello, %s", n)
}
