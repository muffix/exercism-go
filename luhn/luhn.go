// Package luhn contains functions for the Luhn algorithm
package luhn

import (
	"regexp"
	"strconv"
)

// Valid returns whether a string passes the Luhn check
func Valid(number string) bool {
	if !regexp.MustCompile(`^\d[\d\s]+$`).MatchString(number) {
		return false
	}

	return luhn(regexp.MustCompile(`\d`).FindAllString(number, -1))
}

// luhn returns whether a string of digits passes the Luhn check
func luhn(digits []string) bool {
	numDigits := len(digits)
	checksum := 0

	for i := range digits {
		digit, _ := strconv.Atoi(digits[numDigits-i-1])

		if i%2 == 0 {
			checksum += digit
		} else {
			double := 2 * digit
			if double > 9 {
				checksum += double - 9
			} else {
				checksum += double
			}
		}
	}

	return checksum%10 == 0
}
