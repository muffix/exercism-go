// Package scale contains tools for tonic scales
package scale

import "strings"

var sharps = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var flats = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var stepsize = map[string]int{"m": 1, "M": 2, "A": 3}

// Scale generates a new scale
func Scale(tonic, interval string) []string {
	var scale []string

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = sharps
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = flats
	}
	tonic = strings.Title(tonic)

	for i, tone := range scale {
		if tone == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}

	newScale := []string{}

	i := 0
	for _, step := range strings.Split(interval, "") {
		size := stepsize[step]
		newScale = append(newScale, scale[i%len(scale)])
		i += size
	}

	return newScale
}
