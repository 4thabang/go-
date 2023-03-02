package main

import (
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting int", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting string", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})

	t.Run("asserting boolean", func(t *testing.T) {
		AssertEqual(t, true, true)
		AssertNotEqual(t, true, false)
	})
}

type Typer interface {
	string | int | bool
}

func AssertEqual[T Typer](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %+v, want: %+v", got, want)
	}
}

func AssertNotEqual[T Typer](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("%+v should not be equal to: %+v", want, got)
	}
}
