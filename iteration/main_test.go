package main

import (
	"testing"
)

func TestIteration(t *testing.T) {
	expect := "AAAAAA"
	got := Iteration("A", len(expect))
	if expect != got {
		t.Errorf("expected: %q, got: %q", expect, got)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("A", 0)
	}
}

func ExampleIteration() {
	for i := 0; i < 10; i++ {
		Iteration("A", 0)
	}
	// Output: "AAAAAAAAAA"
}
