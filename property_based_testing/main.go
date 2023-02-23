package main

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var (
	allNumerals = RomanNumerals{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

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

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})

			value := allNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total++
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}
