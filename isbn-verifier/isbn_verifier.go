// Package isbn implements an ISBN-10 validator
package isbn

import (
	"regexp"
	"strings"
)

// IsValidISBN returns whether the ISBN is valid
func IsValidISBN(isbn string) bool {
	structureRegex := regexp.MustCompile(`^(\d)\-?(\d{3})\-?(\d{5})\-?([\dX])$`)
	matches := structureRegex.FindStringSubmatch(isbn)
	if len(matches) == 0 {
		return false
	}

	digits := strings.Join(matches[1:], "")
	if len(digits) != 10 {
		return false
	}

	var checksum int

	for i := 0; i < 9; i++ {
		checksum += int(digits[i]-'0') * (10 - i)
	}

	if digits[9] == 'X' {
		checksum += 10
	} else {
		checksum += int(digits[9] - '0')
	}

	return checksum%11 == 0
}
