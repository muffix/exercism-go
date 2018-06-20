//Package pangram contains tools around pangrams
package pangram

import "unicode"

// IsPangram returns whether the given word is a pangram
func IsPangram(s string) bool {
	var letters = make(map[rune]bool)
	for _, c := range s {
		c = unicode.ToLower(c)
		if 'a' <= c && c <= 'z' {
			letters[c] = true
			if len(letters) == 26 {
				return true
			}
		}
	}
	return false
}
