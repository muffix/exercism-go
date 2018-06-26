// Package beer implements a generator for the beer song
package beer

import (
	"fmt"
	"strings"
)

// Song returns the entire beer song
func Song() string {
	song, _ := Verses(99, 0)
	return song
}

// Verses returns the verses of the beer song within the bounds
func Verses(upper, lower int) (string, error) {
	if upper < lower {
		return "", fmt.Errorf("Upper is smaller than lower")
	}

	if lower < 0 || upper > 99 {
		return "", fmt.Errorf("Invalid bounds, must be between 100 and 0")
	}

	var song strings.Builder

	for i := upper; i >= lower; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		song.WriteString(verse)
		song.WriteByte('\n')

	}

	return song.String(), nil
}

// Verse returns verse i of the beer song
func Verse(i int) (string, error) {
	if i < 0 || i > 99 {
		return "", fmt.Errorf("Verse %d doesnt exist. Must be between 0 and 100", i)
	}

	if i == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if i == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if i == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", i, i, i-1), nil
}
