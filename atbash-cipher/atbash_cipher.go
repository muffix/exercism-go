// Package atbash contains an implementatio of the Atbash cipher
package atbash

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[^a-z0-9]`)

// Atbash decodes and encodes a string in the Atbash cipher
func Atbash(in string) string {
	cleaned := re.ReplaceAllString(strings.ToLower(in), "")

	var out strings.Builder

	for i, c := range cleaned {
		if i > 0 && i%5 == 0 {
			out.WriteRune(' ')
		}

		if '0' <= c && c <= '9' {
			out.WriteRune(c)
		} else {
			out.WriteRune('z' - (c - 'a'))
		}
	}

	return out.String()
}
