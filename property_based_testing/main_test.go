package main

import (
	"fmt"
	"testing"
)

type TestingConversion struct {
	Arabic int
	Roman  string
}

var (
	testCases = []TestingConversion{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
)

func TestRomanNumeral(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("%d converts to %s", tc.Arabic, tc.Roman)
		t.Run(testName, func(t *testing.T) {
			got := ConvertToRoman(tc.Arabic)
			if got != tc.Roman {
				t.Errorf("got: %q, want: %q", got, tc.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tc := range testCases[:1] {
		testName := fmt.Sprintf("%d converts to %s", tc.Arabic, tc.Roman)
		t.Run(testName, func(t *testing.T) {
			got := ConvertToArabic(tc.Roman)
			if got != tc.Arabic {
				t.Errorf("got: %d, want: %d", got, tc.Arabic)
			}
		})
	}
}
