// Package proverb implements a proverb generator
package proverb

import (
	"fmt"
)

// Proverb returns the proverb
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	verses := make([]string, len(rhyme))

	for i := 0; i < len(rhyme)-1; i++ {
		verses[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	verses[len(verses)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return verses
}
