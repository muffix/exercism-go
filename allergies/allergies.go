// Package allergies contains tools to determine allergies
package allergies

var allergyScores = map[string]uint{
	"eggs":         0,
	"peanuts":      1,
	"shellfish":    2,
	"strawberries": 3,
	"tomatoes":     4,
	"chocolate":    5,
	"pollen":       6,
	"cats":         7,
}

var allergyOrder = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns all substances that someone with the score is allergic to
func Allergies(score uint) []string {
	var results []string

	for _, substance := range allergyOrder {
		if AllergicTo(score, substance) {
			results = append(results, substance)
		}
	}
	return results
}

// AllergicTo returns whether someone with the score is allergic to the substance
func AllergicTo(score uint, substance string) bool {
	if score&(1<<allergyScores[substance]) > 0 {
		return true
	}
	return false
}
