package main_test

import (
	go_specs_greet "go-tests/acceptance-tests"
	"go-tests/acceptance-tests/specifications"
	"testing"
)

func TestGreetServer(t *testing.T) {
	driver := go_specs_greet.Driver{
		BaseURL: "http://localhost:8080",
	}
	specifications.GreetSpecification(t, driver)
}
