// Package phonenumber contains functions to format phone numbers
package phonenumber

import (
	"fmt"
	"regexp"
)

var digitsRE = regexp.MustCompile(`[^\d]`)
var areaCodeRE = regexp.MustCompile(`^1?([2-9]\d{2})[2-9]\d{6}$`)
var phoneValidationRE = regexp.MustCompile(`^1?((?:[2-9]\d{2}){2}\d{4})$`)

// Number returns the digits-only string of a phone number
func Number(number string) (string, error) {
	digits := digitsRE.ReplaceAllString(number, "")
	matches := phoneValidationRE.FindStringSubmatch(digits)

	if len(matches) == 0 {
		return "", fmt.Errorf("Number has an invalid format")
	}

	return matches[len(matches)-1], nil
}

// AreaCode returns the area code of a phone number
func AreaCode(number string) (string, error) {
	digits := digitsRE.ReplaceAllString(number, "")
	matches := areaCodeRE.FindStringSubmatch(digits)

	if len(matches) > 0 {
		return matches[len(matches)-1], nil
	}

	return "", fmt.Errorf("Number has an invalid format")
}

// Format returns the number formatted to (NXX)-NXX-XXXX
func Format(number string) (string, error) {
	number, err := Number(number)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
