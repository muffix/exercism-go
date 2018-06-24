// Package romannumerals contains a converter for integers to Roman numerals
package romannumerals

import (
	"fmt"
	"strings"
)

type numeral struct {
	symbol string
	value  int
}

var numerals = []numeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts an integer to its roman numeral representation
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3000 {
		return "", fmt.Errorf("Number must be between 1 and 3000, got: %d", num)
	}

	var roman strings.Builder

	for _, numeral := range numerals {
		for num >= numeral.value {
			roman.WriteString(numeral.symbol)
			num -= numeral.value
		}
		if num == 0 {
			break
		}
	}

	return roman.String(), nil
}
