// Package railfence contains tools to decode/encode strings with a rail fence cipher
package railfence

const (
	decodeAction = iota
	encodeAction
)

// Encode encodes a message using a rail fence cipher with n rails
func Encode(plaintext string, n int) string {
	return process(plaintext, n, encodeAction)
}

// Decode decodes a message using a rail fence cipher with n rails
func Decode(ciphertext string, n int) string {
	return process(ciphertext, n, decodeAction)
}

func process(in string, numRails, action int) string {
	inRunes := []rune(in)
	out := []rune(in)

	var i, column int
	var inPos, outPos *int

	if action == encodeAction {
		inPos = &column
		outPos = &i
	} else {
		inPos = &i
		outPos = &column
	}

	longestStep := 2 * (numRails - 1)
	for rail := 0; rail < numRails; rail++ {
		nextStep := 2 * rail
		for column = rail; column < len(in); column += nextStep {
			out[*outPos] = inRunes[*inPos]
			i++
			// The first and last rail always have the same step size
			if nextStep == longestStep {
				continue
			}
			// The other rails alternate between two step sizes
			nextStep = longestStep - nextStep
		}
	}
	return string(out)
}
