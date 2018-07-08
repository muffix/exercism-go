// Package rotationalcipher implements a rotational cipher
package rotationalcipher

import (
	"strings"
)

// RotationalCipher returns the rotated text
func RotationalCipher(text string, key int) string {
	var out strings.Builder

	for i := range text {
		char := text[i]
		switch {

		case char >= 'a' && char <= 'z':
			out.WriteByte((char-'a'+byte(key))%26 + 'a')
		case char >= 'A' && char <= 'Z':
			out.WriteByte((char-'A'+byte(key))%26 + 'A')
		default:
			out.WriteByte(char)
		}
	}

	return out.String()
}
