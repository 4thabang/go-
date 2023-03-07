package main

import (
	"fmt"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

type driver string

func (d driver) Greet() (string, error) {
	if d == "" {
		return "", fmt.Errorf("empty string")
	}
	return fmt.Sprintf("Hello, %s", d), nil
}

func TestGreeter(t *testing.T) {
	d := driver("world")
	specifications.GreetSpecification(t, d)
}
