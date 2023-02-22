package main

import (
	"strings"
)

func ConvertToRoman(arabic int) string {
	var str strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 9:
			str.WriteString("X")
			arabic -= 10
		case arabic > 8:
			str.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			str.WriteString("V")
			arabic -= 5
		case arabic > 3:
			str.WriteString("IV")
			arabic -= 4
		default:
			str.WriteString("I")
			arabic--
		}
	}

	return str.String()
}
