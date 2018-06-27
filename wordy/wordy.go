// Package wordy implements a solver for simple math word problems
package wordy

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(?:What is (-?\d+))|(?:(multiplied|divided|plus|minus) (?:by )?(-?\d+))`)

// Answer returns the answer to a mathematical question
func Answer(q string) (int, bool) {
	matches := re.FindAllStringSubmatch(q, -1)

	if len(matches) < 2 {
		return 0, false
	}

	result, err := strconv.Atoi(matches[0][1])

	if err != nil {
		return 0, false
	}

	for i := 1; i < len(matches); i++ {
		operand, err := strconv.Atoi(matches[i][3])

		if err != nil {
			return 0, false
		}

		switch matches[i][2] {
		case "multiplied":
			result *= operand
		case "divided":
			result /= operand
		case "plus":
			result += operand
		case "minus":
			result -= operand
		}
	}

	return result, true
}
