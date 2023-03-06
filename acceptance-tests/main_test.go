package main

import (
	"fmt"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

type some string

func (s some) Greet() (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty string")
	}
	return fmt.Sprintf("Hello, %s", s), nil
}

func TestGreeter(t *testing.T) {
	s := some("")
	specifications.GreetSpecification(t, s)
}
