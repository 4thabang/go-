package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		testName       string
		name, language string
		want           string
	}{
		{
			testName: "Chris is English",
			name:     "Chris",
			language: "English",
			want:     "Hello, Chris",
		},
		{
			testName: "Empty everything",
			name:     "",
			language: "",
			want:     "Hello, World",
		},
		{
			testName: "Elodie is Spanish",
			name:     "Elodie",
			language: "Spanish",
			want:     "Hola, Elodie",
		},
		{
			testName: "Genevive is French",
			name:     "enevive",
			language: "French",
			want:     "Bonjour, Genevive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got := Hello(tt.name, tt.language)
			want := tt.want
			assertCorrectMessage(t, got, want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
