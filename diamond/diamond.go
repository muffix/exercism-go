// Package diamond contains tools to generate pretty diamonds from letters
package diamond

import (
	"fmt"
	"strings"
)

// Gen generates a diamond with b at the maximum width
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", fmt.Errorf("Expected character between A and Z, got: %s", string(b))
	}

	var rows []string
	current := byte('A')

	for current <= b {
		skip := strings.Repeat(" ", int(b-current))

		if current == 'A' {
			rows = append(rows, fmt.Sprintln(skip+string(current)+skip))
		} else {
			middle := strings.Repeat(" ", int(2*(current-'A')-1))
			rows = append(rows, fmt.Sprintln(skip+string(current)+middle+string(current)+skip))
		}

		current++
	}

	var diamond strings.Builder

	diamond.WriteString(strings.Join(rows, ""))

	for i := len(rows) - 2; i >= 0; i-- {
		diamond.WriteString(rows[i])
	}

	return diamond.String(), nil
}
