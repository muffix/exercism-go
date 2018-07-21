// Package rectangles detects rectangles
package rectangles

import (
	"regexp"
	"strings"
)

type field []string
type corners [4][2]int

var rowRe = regexp.MustCompile(`^\+[-\+]*\+$`)
var colRe = regexp.MustCompile(`^\+[\|\+]*\+$`)

// Count returns the number of rectangles in the input
func Count(input field) int {
	var allCorners [][2]int
	for i, row := range input {
		for j, col := range row {
			if col == '+' {
				allCorners = append(allCorners, [2]int{i, j})
			}
		}
	}

	combinationsCh := make(chan [4]int)
	go combinations(len(allCorners), combinationsCh)

	count := 0
	for combi := range combinationsCh {
		var c corners

		for i, k := range combi {
			c[i] = allCorners[k]
		}

		if input.isRectangle(c) {
			count++
		}
	}

	return count
}

func (f field) isRectangle(c corners) bool {
	tl, tr, bl, br := c[0], c[1], c[2], c[3]

	if tl[0] != tr[0] || bl[0] != br[0] {
		return false
	}

	if tl[1] != bl[1] || tr[1] != br[1] {
		return false
	}

	topRow, bottomRow := f[tl[0]][tl[1]:tr[1]+1], f[bl[0]][bl[1]:br[1]+1]
	var firstCol, lastCol strings.Builder

	if !rowRe.MatchString(topRow) || !rowRe.MatchString(bottomRow) {
		return false
	}

	for i := tl[0]; i <= bl[0]; i++ {
		firstCol.WriteByte(f[i][tl[1]])
		lastCol.WriteByte(f[i][tr[1]])
	}

	if !colRe.MatchString(firstCol.String()) || !colRe.MatchString(lastCol.String()) {
		return false
	}

	return true
}

func combinations(n int, ch chan<- [4]int) {
	if n < 4 {
		close(ch)
		return
	}

	var combination [4]int
	for i := 0; i < 4; i++ {
		combination[i] = i
	}

	ch <- combination

	for {
		var i int
		for i = 4 - 1; i >= 0 && combination[i] == n-4+i; i-- {
		}
		if i < 0 {
			break
		}

		combination[i]++

		i++
		for ; i < 4; i++ {
			combination[i] = combination[i-1] + 1
		}
		ch <- combination
	}
	close(ch)
}
