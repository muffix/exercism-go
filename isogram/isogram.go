// Package isogram contains functions to determine if a word is an isogram
package isogram

import "unicode"

// IsIsogram determines if a word is an isogram
func IsIsogram(maybeIsogram string) bool {
	seen := map[rune]bool{}

	for _, letter := range maybeIsogram {
		if unicode.IsLetter(letter) {
			lowercaseLetter := unicode.ToLower(letter)
			if seen[lowercaseLetter] {
				return false
			}
			seen[lowercaseLetter] = true
		}
	}

	return true
}
