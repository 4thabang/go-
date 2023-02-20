package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("multiplication", func(t *testing.T) {
		expect := 10
		got := Adder(5, 2)
		assertAdder(t, got, expect)
	})
}

func assertAdder(t testing.TB, got, expect int) {
	t.Helper()
	if got != expect {
		t.Errorf("expected: '%d', got: '%d'", expect, got)
	}
}

func ExampleAdder() {
	sum := Adder(10, 20)
	fmt.Println(sum)
	// Output: 200
}
