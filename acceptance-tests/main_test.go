package main

import (
	"fmt"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

type greeting string

func (s greeting) Greet() (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty string")
	}
	return fmt.Sprintf("Hello, %s", s), nil
}

func TestGreeter(t *testing.T) {
	s := greeting("world")
	specifications.GreetSpecification(t, s)
}
