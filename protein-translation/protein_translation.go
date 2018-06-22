// Package protein contains tools around pProteins
package protein

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

// FromCodon returns the protein for a given codon
func FromCodon(codon string) string {
	return codons[codon]
}

// FromRNA returns a slice of proteins for the given RNA
func FromRNA(rna string) []string {
	var proteins []string

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein := FromCodon(codon)

		if protein == "STOP" {
			return proteins
		}

		proteins = append(proteins, protein)
	}
	return proteins
}
