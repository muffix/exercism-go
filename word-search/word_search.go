// Package wordsearch implements a simple word search
package wordsearch

import (
	"fmt"
	"strings"
)

type coordinates struct {
	x, y int
}

type path []coordinates

// Solve returns the start and end coordinates in the puzzle for each of the requested words
func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	results := map[string][2][2]int{}

	for _, word := range words {
		start, end, err := findWord(word, puzzle)

		if err != nil {
			return nil, err
		}

		results[word] = [2][2]int{{start.x, start.y}, {end.x, end.y}}
	}

	if len(results) == 0 {
		return results, fmt.Errorf("No results")
	}

	return results, nil
}

// findWord finds a single word in the puzzle and returns its start and end coordinates if found
func findWord(word string, puzzle []string) (coordinates, coordinates, error) {
	rows := len(puzzle)
	cols := len(puzzle[0])

	for r, row := range puzzle {
		for c, char := range row {
			if byte(char) == word[0] {
				possiblePaths := paths(coordinates{c, r}, len(word), rows, cols)
				for _, path := range possiblePaths {
					if path.word(puzzle) == word {
						return path[0], path[len(path)-1], nil
					}
				}
			}
		}
	}

	return coordinates{}, coordinates{}, fmt.Errorf("Word %s not found", word)
}

// word returns the word that this path spells in the puzzle
func (p path) word(puzzle []string) string {
	var word strings.Builder

	for _, c := range p {
		word.WriteByte(puzzle[c.y][c.x])
	}

	return word.String()
}

// paths constructs all valid paths of the given length from the origin
func paths(origin coordinates, length, rows, cols int) []path {
	var paths []path
	length--

	for i := -length; i <= length; i += length {
		for j := -length; j <= length; j += length {
			if i == 0 && j == 0 {
				continue
			}

			candidate := coordinates{origin.x + i, origin.y + j}
			if candidate.isInBounds(rows, cols) {
				paths = append(paths, fillPath(origin, candidate, i/length, j/length))
			}
		}
	}

	return paths
}

// isInBounds returns whether the coordinates are in the given bounds
func (c coordinates) isInBounds(rows, cols int) bool {
	return c.x >= 0 && c.x < cols && c.y >= 0 && c.y < rows
}

// fillPath returns the full path with all intermediate steps between the two coordinates
func fillPath(from, to coordinates, dx, dy int) path {
	path := []coordinates{from}

	next := from
	for next != to {
		next = coordinates{next.x + dx, next.y + dy}
		path = append(path, next)
	}

	return path
}
