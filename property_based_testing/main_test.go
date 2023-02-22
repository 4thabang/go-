package main

import (
	"testing"
)

func TestRomanNumeral(t *testing.T) {
	testCases := []struct {
		TestName string
		Arabic   int
		Want     string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 converts to III", 3, "III"},
		{"4 converts to IV", 4, "IV"},
		{"5 converts to V", 5, "V"},
		{"6 converts to VI", 6, "VI"},
		{"7 converts to VII", 7, "VII"},
		{"8 converts to VIII", 8, "VIII"},
		{"9 converts to IX", 9, "IX"},
		{"10 converts to X", 10, "X"},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got := ConvertToRoman(tc.Arabic)

			if got != tc.Want {
				t.Errorf("got: %q, want: %q", got, tc.Want)
			}
		})
	}
}
