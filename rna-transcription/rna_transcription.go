// Package strand contains tools around strands
package strand

import "strings"

var dnaToRNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns the RNA transcription of the DNA strand
func ToRNA(dna string) string {
	var rna strings.Builder

	for _, nucleotide := range dna {
		rna.WriteRune(dnaToRNA[nucleotide])
	}

	return rna.String()
}
