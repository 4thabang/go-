package main

import (
	"fmt"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

type system string

func (s system) Greet() (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty string")
	}
	return fmt.Sprintf("Hello, %s", s), nil
}

func TestGreeter(t *testing.T) {
	s := system("world")
	specifications.GreetSpecification(t, s)
}
