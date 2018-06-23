// Package wordcount contains tools around counting words
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is the datastructure representing the frequency of words
type Frequency map[string]int

var wordRegex = regexp.MustCompile(`[\w']+`)

// WordCount returns the frequency per word in the phrase
func WordCount(phrase string) Frequency {
	frequency := Frequency{}

	words := wordRegex.FindAllString(phrase, -1)

	for _, word := range words {
		frequency[strings.Trim(strings.ToLower(word), "'")]++
	}

	return frequency
}
