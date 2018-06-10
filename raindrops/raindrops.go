// Package raindrops contains functions for handling raindrops
package raindrops

import (
	"fmt"
	"strings"
)

// Convert converts the number of raindrops into their string representation
func Convert(num int) string {
	var convertedBuilder strings.Builder

	if num%3 == 0 {
		convertedBuilder.WriteString("Pling")
	}
	if num%5 == 0 {
		convertedBuilder.WriteString("Plang")
	}
	if num%7 == 0 {
		convertedBuilder.WriteString("Plong")
	}

	converted := convertedBuilder.String()
	if converted == "" {
		return fmt.Sprintf("%d", num)
	}

	return converted
}
