// Package etl contains transformations of Scrabble score notations
package etl

import (
	"strings"
)

// Transform transforms the old Scrabble scoring format into the new one
func Transform(oldScores map[int][]string) map[string]int {
	newScores := make(map[string]int)

	for value, letters := range oldScores {
		for _, letter := range letters {
			newScores[strings.ToLower(letter)] = value
		}
	}

	return newScores
}
