package cipher

import (
	"regexp"
	"strings"
)

const (
	vigenereEncodeOp = 1
	vigenereDecodeOp = -1
)

// Caesar represents the Caesar cipher
type Caesar struct{}

// Shift represents a Caesar cipher with a flexible shift distance
type Shift struct {
	key int
}

// Vigenere represents a simple Vigenere cipher
type Vigenere struct {
	key string
}

var notLowercaseLetterRE = regexp.MustCompile(`[^a-z]`)
var vigenereValidationRE = regexp.MustCompile(`^a*$|[^a-z]`)

// NewCaesar returns a new Caesar cipher
func NewCaesar() *Caesar {
	return &Caesar{}
}

// NewShift returns a new Caesar cipher with a flexible shift distance
func NewShift(key int) *Shift {
	if key == 0 || key > 25 || key < -25 {
		return nil
	}
	return &Shift{key}
}

// NewVigenere returns a new Vigenere cipher with the given key
func NewVigenere(key string) *Vigenere {
	if ok := validateVigenereKey(key); !ok {
		return nil
	}
	return &Vigenere{key}
}

// Encode returns the encoded message
func (c Caesar) Encode(plaintext string) string {
	return shiftString(convertText(plaintext), 3)
}

// Decode returns the decoded message
func (c Caesar) Decode(ciphertext string) string {
	return shiftString(convertText(ciphertext), -3)
}

// Encode returns the encoded message
func (s Shift) Encode(plaintext string) string {
	return shiftString(convertText(plaintext), s.key)
}

// Decode returns the decoded message
func (s Shift) Decode(ciphertext string) string {
	return shiftString(convertText(ciphertext), -s.key)
}

// Encode returns the encoded message
func (v Vigenere) Encode(plaintext string) string {
	return vigenere(convertText(plaintext), v.key, vigenereEncodeOp)
}

// Decode returns the decoded message
func (v Vigenere) Decode(ciphertext string) string {
	return vigenere(convertText(ciphertext), v.key, vigenereDecodeOp)
}

func shiftRune(r rune, offset int) rune {
	r += rune(offset)

	if r > 'z' {
		r -= 'z' - 'a' + 1
	} else if r < 'a' {
		r += 'z' - 'a' + 1
	}

	return r
}

func shiftString(s string, offset int) string {
	var b strings.Builder

	for _, r := range s {
		b.WriteRune(shiftRune(r, offset))
	}
	return b.String()
}

func convertText(text string) string {
	return notLowercaseLetterRE.ReplaceAllString(strings.ToLower(text), "")
}

func validateVigenereKey(key string) bool {
	return !vigenereValidationRE.MatchString(key)
}

func vigenere(text, key string, op int) string {
	var converted strings.Builder

	keyLength := len(key)

	for i, r := range text {
		offset := ((key[i%keyLength] - 'a') % 26)
		converted.WriteRune(shiftRune(r, op*int(offset)))
	}

	return converted.String()
}
