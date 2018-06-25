// Package transpose contains tools to transpose a matrix of strings
package transpose

// Transpose returns a new array of strings with the transposed input
func Transpose(m []string) []string {
	output := []string{}

	for r, row := range m {
		for c, char := range row {
			for len(output) <= c {
				output = append(output, "")
			}

			for len(output[c]) < r {
				output[c] += " "
			}

			output[c] += string(char)
		}
	}

	return output
}
