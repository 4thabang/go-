package main

import (
	"strings"
)

type RomanNumerals struct {
	Value  int
	Symbol string
}

var (
	allNumerals = []RomanNumerals{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

func ConvertToRoman(arabic int) string {
	var str strings.Builder

	for _, numeral := range allNumerals {
		for arabic >= numeral.Value {
			str.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return str.String()
}
