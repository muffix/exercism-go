// Package secret contains tools to construct a secret handshake
package secret

var actions = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake returns the secret handshake for the given code
func Handshake(code uint) []string {
	reverse := code&(1<<uint(len(actions))) > 0

	var sequence []string

	for i := 0; i < len(actions); i++ {
		exp := i
		if reverse {
			exp = len(actions) - i - 1
		}

		if code&(1<<uint(exp)) > 0 {
			sequence = append(sequence, actions[exp])
		}
	}

	return sequence
}
