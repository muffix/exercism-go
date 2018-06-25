// Package piglatin contains tools to translate a string to Pig Latiin
package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

var startsWithVowel = regexp.MustCompile(`^(xr|yt|[aeiou]+)(.*)$`)
var yAfterConsonants = regexp.MustCompile(`^([^aeiouy]+)((y).*)`)
var quAfterConsonant = regexp.MustCompile(`^([^aeiou]*qu)(.*)`)
var startsWithConsonant = regexp.MustCompile(`^([^aeiou]*)(.*)$`)

// Sentence returns the sentence translated to Pig Latin
func Sentence(s string) string {
	var latin strings.Builder
	for _, word := range strings.Split(s, " ") {
		latin.WriteString(wordToPigLatin(word) + " ")
	}

	return strings.TrimSpace(latin.String())
}

func wordToPigLatin(s string) string {
	if ok := startsWithVowel.MatchString(s); ok {
		return fmt.Sprintf("%say", s)
	}

	matches := yAfterConsonants.FindStringSubmatch(s)

	if len(matches) == 0 {
		matches = quAfterConsonant.FindStringSubmatch(s)
		if len(matches) == 0 {
			matches = startsWithConsonant.FindStringSubmatch(s)
		}
	}

	return fmt.Sprintf("%s%say", matches[2], matches[1])
}
