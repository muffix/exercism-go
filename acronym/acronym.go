// Package acronym contains methods to generate acronyms
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns an uppercase acronym for a given string
func Abbreviate(s string) string {
	words := regexp.MustCompile(`\p{L}+`).FindAllString(s, -1)

	var abbreviation strings.Builder

	for _, word := range words {
		abbreviation.WriteString(string([]rune(word)[0]))
	}

	return strings.ToUpper(abbreviation.String())
}
