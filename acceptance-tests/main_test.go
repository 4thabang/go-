package main

import (
	"fmt"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

type driver string

func (s driver) Greet() (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty string")
	}
	return fmt.Sprintf("Hello, %s", s), nil
}

func TestGreeter(t *testing.T) {
	s := driver("world")
	specifications.GreetSpecification(t, s)
}
