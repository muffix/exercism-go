// Package minesweeper implements minesweeper
package minesweeper

import (
	"fmt"
	"regexp"
)

// Count fills the board with numbers according to the Minesweeper rules
func (b Board) Count() error {
	if ok := b.validate(); !ok {
		return fmt.Errorf("Invalid board")
	}

	rows, cols := len(b), len(b[0])

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if b[i][j] != '*' {
				b.fillNumber(i, j, rows, cols)
			}
		}
	}

	return nil
}

func (b Board) validate() bool {
	outer := regexp.MustCompile(`^\+-*\+$`)

	headValid := outer.Match(b[0])

	if !headValid {
		return false
	}

	footValid := outer.Match(b[len(b)-1])

	if !footValid {
		return false
	}

	inner := regexp.MustCompile(fmt.Sprintf("^\\|(?:[* ]){%d}\\|$", len(b[0])-2))

	for i := 1; i < len(b)-1; i++ {
		if !inner.Match(b[i]) {
			return false
		}
	}

	return true
}

func (b Board) fillNumber(i, j, rows, cols int) {
	var count byte

	for _, n := range neighboursinBounds(i, j, rows, cols) {
		if b[n[0]][n[1]] == '*' {
			count++
		}
	}

	if count != 0 {
		b[i][j] = '0' + count
	}
}

func neighboursinBounds(i, j, rows, cols int) [][2]int {
	neighbours := [][2]int{}

	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if k != 0 || l != 0 {
				if i+k > 0 && i+k < rows-1 && j+l > 0 && j+l < cols-1 {
					neighbours = append(neighbours, [2]int{i + k, j + l})
				}
			}
		}
	}

	return neighbours
}
