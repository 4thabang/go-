package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buffer bytes.Buffer
	name := Name("Thabang")

	name.Greet(&buffer)

	got := buffer.String()
	want := "Hello, Thabang"

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
